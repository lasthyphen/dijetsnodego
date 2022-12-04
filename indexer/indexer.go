// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package indexer

import (
	"fmt"
	"io"
	"math"
	"sync"

	"github.com/gorilla/rpc/v2"

	"go.uber.org/zap"

	"github.com/lasthyphen/dijetsnodego/api/server"
	"github.com/lasthyphen/dijetsnodego/chains"
	"github.com/lasthyphen/dijetsnodego/codec"
	"github.com/lasthyphen/dijetsnodego/codec/linearcodec"
	"github.com/lasthyphen/dijetsnodego/database"
	"github.com/lasthyphen/dijetsnodego/database/prefixdb"
	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/snow"
	"github.com/lasthyphen/dijetsnodego/snow/engine/avalanche"
	"github.com/lasthyphen/dijetsnodego/snow/engine/common"
	"github.com/lasthyphen/dijetsnodego/snow/engine/snowman"
	"github.com/lasthyphen/dijetsnodego/utils/constants"
	"github.com/lasthyphen/dijetsnodego/utils/hashing"
	"github.com/lasthyphen/dijetsnodego/utils/json"
	"github.com/lasthyphen/dijetsnodego/utils/logging"
	"github.com/lasthyphen/dijetsnodego/utils/timer/mockable"
	"github.com/lasthyphen/dijetsnodego/utils/wrappers"
)

const (
	indexNamePrefix = "index-"
	codecVersion    = uint16(0)
	// Max size, in bytes, of something serialized by this indexer
	// Assumes no containers are larger than math.MaxUint32
	// wrappers.IntLen accounts for the size of the container bytes
	// wrappers.LongLen accounts for the timestamp of the container
	// hashing.HashLen accounts for the container ID
	// wrappers.ShortLen accounts for the codec version
	codecMaxSize = int(constants.DefaultMaxMessageSize) + wrappers.IntLen + wrappers.LongLen + hashing.HashLen + wrappers.ShortLen
)

var (
	txPrefix                = byte(0x01)
	vtxPrefix               = byte(0x02)
	blockPrefix             = byte(0x03)
	isIncompletePrefix      = byte(0x04)
	previouslyIndexedPrefix = byte(0x05)
	hasRunKey               = []byte{0x07}

	_ Indexer = &indexer{}
)

// Config for an indexer
type Config struct {
	DB                     database.Database
	Log                    logging.Logger
	IndexingEnabled        bool
	AllowIncompleteIndex   bool
	DecisionAcceptorGroup  snow.AcceptorGroup
	ConsensusAcceptorGroup snow.AcceptorGroup
	APIServer              server.PathAdder
	ShutdownF              func()
}

// Indexer causes accepted containers for a given chain
// to be indexed by their ID and by the order in which
// they were accepted by this node.
// Indexer is threadsafe.
type Indexer interface {
	chains.Registrant
	// Close will do nothing and return nil after the first call
	io.Closer
}

// NewIndexer returns a new Indexer and registers a new endpoint on the given API server.
func NewIndexer(config Config) (Indexer, error) {
	indexer := &indexer{
		codec:                  codec.NewManager(codecMaxSize),
		log:                    config.Log,
		db:                     config.DB,
		allowIncompleteIndex:   config.AllowIncompleteIndex,
		indexingEnabled:        config.IndexingEnabled,
		decisionAcceptorGroup:  config.DecisionAcceptorGroup,
		consensusAcceptorGroup: config.ConsensusAcceptorGroup,
		txIndices:              map[ids.ID]Index{},
		vtxIndices:             map[ids.ID]Index{},
		blockIndices:           map[ids.ID]Index{},
		pathAdder:              config.APIServer,
		shutdownF:              config.ShutdownF,
	}

	if err := indexer.codec.RegisterCodec(
		codecVersion,
		linearcodec.NewCustomMaxLength(math.MaxUint32),
	); err != nil {
		return nil, fmt.Errorf("couldn't register codec: %w", err)
	}
	hasRun, err := indexer.hasRun()
	if err != nil {
		return nil, err
	}
	indexer.hasRunBefore = hasRun
	return indexer, indexer.markHasRun()
}

type indexer struct {
	codec  codec.Manager
	clock  mockable.Clock
	lock   sync.RWMutex
	log    logging.Logger
	db     database.Database
	closed bool

	// Called in a goroutine on shutdown
	shutdownF func()

	// true if this is not the first run using this database
	hasRunBefore bool

	// Used to add API endpoint for new indices
	pathAdder server.PathAdder

	// If true, allow running in such a way that could allow the creation
	// of an index which could be missing accepted containers.
	allowIncompleteIndex bool

	// If false, don't create index for a chain when RegisterChain is called
	indexingEnabled bool

	// Chain ID --> index of blocks of that chain (if applicable)
	blockIndices map[ids.ID]Index
	// Chain ID --> index of vertices of that chain (if applicable)
	vtxIndices map[ids.ID]Index
	// Chain ID --> index of txs of that chain (if applicable)
	txIndices map[ids.ID]Index

	// Notifies of newly accepted transactions
	decisionAcceptorGroup snow.AcceptorGroup
	// Notifies of newly accepted blocks and vertices
	consensusAcceptorGroup snow.AcceptorGroup
}

// Assumes [engine]'s context lock is not held
func (i *indexer) RegisterChain(name string, engine common.Engine) {
	i.lock.Lock()
	defer i.lock.Unlock()

	ctx := engine.Context()
	if i.closed {
		i.log.Debug("not registering chain to indexer",
			zap.String("reason", "indexer is closed"),
			zap.String("chainName", name),
		)
		return
	} else if ctx.SubnetID != constants.PrimaryNetworkID {
		i.log.Debug("not registering chain to indexer",
			zap.String("reason", "not in the primary network"),
			zap.String("chainName", name),
		)
		return
	}

	chainID := ctx.ChainID
	if i.blockIndices[chainID] != nil || i.txIndices[chainID] != nil || i.vtxIndices[chainID] != nil {
		i.log.Warn("chain is already being indexed",
			zap.Stringer("chainID", chainID),
		)
		return
	}

	// If the index is incomplete, make sure that's OK. Otherwise, cause node to die.
	isIncomplete, err := i.isIncomplete(chainID)
	if err != nil {
		i.log.Error("couldn't get whether chain is incomplete",
			zap.String("chainName", name),
			zap.Error(err),
		)
		if err := i.close(); err != nil {
			i.log.Error("failed to close indexer",
				zap.Error(err),
			)
		}
		return
	}

	// See if this chain was indexed in a previous run
	previouslyIndexed, err := i.previouslyIndexed(chainID)
	if err != nil {
		i.log.Error("couldn't get whether chain was previously indexed",
			zap.String("chainName", name),
			zap.Error(err),
		)
		if err := i.close(); err != nil {
			i.log.Error("failed to close indexer",
				zap.Error(err),
			)
		}
		return
	}

	if !i.indexingEnabled { // Indexing is disabled
		if previouslyIndexed && !i.allowIncompleteIndex {
			// We indexed this chain in a previous run but not in this run.
			// This would create an incomplete index, which is not allowed, so exit.
			i.log.Fatal("running would cause index to become incomplete but incomplete indices are disabled",
				zap.String("chainName", name),
			)
			if err := i.close(); err != nil {
				i.log.Error("failed to close indexer",
					zap.Error(err),
				)
			}
			return
		}

		// Creating an incomplete index is allowed. Mark index as incomplete.
		err := i.markIncomplete(chainID)
		if err == nil {
			return
		}
		i.log.Fatal("couldn't mark chain as incomplete",
			zap.String("chainName", name),
			zap.Error(err),
		)
		if err := i.close(); err != nil {
			i.log.Error("failed to close indexer",
				zap.Error(err),
			)
		}
		return
	}

	if !i.allowIncompleteIndex && isIncomplete && (previouslyIndexed || i.hasRunBefore) {
		i.log.Fatal("index is incomplete but incomplete indices are disabled. Shutting down",
			zap.String("chainName", name),
		)
		if err := i.close(); err != nil {
			i.log.Error("failed to close indexer",
				zap.Error(err),
			)
		}
		return
	}

	// Mark that in this run, this chain was indexed
	if err := i.markPreviouslyIndexed(chainID); err != nil {
		i.log.Error("couldn't mark chain as indexed",
			zap.String("chainName", name),
			zap.Error(err),
		)
		if err := i.close(); err != nil {
			i.log.Error("failed to close indexer",
				zap.Error(err),
			)
		}
		return
	}

	switch engine.(type) {
	case snowman.Engine:
		index, err := i.registerChainHelper(chainID, blockPrefix, name, "block", i.consensusAcceptorGroup)
		if err != nil {
			i.log.Fatal("failed to create block index",
				zap.String("chainName", name),
				zap.Error(err),
			)
			if err := i.close(); err != nil {
				i.log.Error("failed to close indexer",
					zap.Error(err),
				)
			}
			return
		}
		i.blockIndices[chainID] = index
	case avalanche.Engine:
		vtxIndex, err := i.registerChainHelper(chainID, vtxPrefix, name, "vtx", i.consensusAcceptorGroup)
		if err != nil {
			i.log.Fatal("couldn't create vertex index",
				zap.String("chainName", name),
				zap.Error(err),
			)
			if err := i.close(); err != nil {
				i.log.Error("failed to close indexer",
					zap.Error(err),
				)
			}
			return
		}
		i.vtxIndices[chainID] = vtxIndex

		txIndex, err := i.registerChainHelper(chainID, txPrefix, name, "tx", i.decisionAcceptorGroup)
		if err != nil {
			i.log.Fatal("couldn't create tx index for",
				zap.String("chainName", name),
				zap.Error(err),
			)
			if err := i.close(); err != nil {
				i.log.Error("failed to close indexer:",
					zap.Error(err),
				)
			}
			return
		}
		i.txIndices[chainID] = txIndex
	default:
		engineType := fmt.Sprintf("%T", engine)
		i.log.Error("got unexpected engine type",
			zap.String("engineType", engineType),
		)
		if err := i.close(); err != nil {
			i.log.Error("failed to close indexer",
				zap.Error(err),
			)
		}
		return
	}
}

func (i *indexer) registerChainHelper(
	chainID ids.ID,
	prefixEnd byte,
	name, endpoint string,
	acceptorGroup snow.AcceptorGroup,
) (Index, error) {
	prefix := make([]byte, hashing.HashLen+wrappers.ByteLen)
	copy(prefix, chainID[:])
	prefix[hashing.HashLen] = prefixEnd
	indexDB := prefixdb.New(prefix, i.db)
	index, err := newIndex(indexDB, i.log, i.codec, i.clock)
	if err != nil {
		_ = indexDB.Close()
		return nil, err
	}

	// Register index to learn about new accepted vertices
	if err := acceptorGroup.RegisterAcceptor(chainID, fmt.Sprintf("%s%s", indexNamePrefix, chainID), index, true); err != nil {
		_ = index.Close()
		return nil, err
	}

	// Create an API endpoint for this index
	apiServer := rpc.NewServer()
	codec := json.NewCodec()
	apiServer.RegisterCodec(codec, "application/json")
	apiServer.RegisterCodec(codec, "application/json;charset=UTF-8")
	if err := apiServer.RegisterService(&service{Index: index}, "index"); err != nil {
		_ = index.Close()
		return nil, err
	}
	handler := &common.HTTPHandler{LockOptions: common.NoLock, Handler: apiServer}
	if err := i.pathAdder.AddRoute(handler, &sync.RWMutex{}, "index/"+name, "/"+endpoint); err != nil {
		_ = index.Close()
		return nil, err
	}
	return index, nil
}

// Close this indexer. Stops indexing all chains.
// Closes [i.db]. Assumes Close is only called after
// the node is done making decisions.
// Calling Close after it has been called does nothing.
func (i *indexer) Close() error {
	i.lock.Lock()
	defer i.lock.Unlock()

	return i.close()
}

func (i *indexer) close() error {
	if i.closed {
		return nil
	}
	i.closed = true

	errs := &wrappers.Errs{}
	for chainID, txIndex := range i.txIndices {
		errs.Add(
			txIndex.Close(),
			i.decisionAcceptorGroup.DeregisterAcceptor(chainID, fmt.Sprintf("%s%s", indexNamePrefix, chainID)),
		)
	}
	for chainID, vtxIndex := range i.vtxIndices {
		errs.Add(
			vtxIndex.Close(),
			i.consensusAcceptorGroup.DeregisterAcceptor(chainID, fmt.Sprintf("%s%s", indexNamePrefix, chainID)),
		)
	}
	for chainID, blockIndex := range i.blockIndices {
		errs.Add(
			blockIndex.Close(),
			i.consensusAcceptorGroup.DeregisterAcceptor(chainID, fmt.Sprintf("%s%s", indexNamePrefix, chainID)),
		)
	}
	errs.Add(i.db.Close())

	go i.shutdownF()
	return errs.Err
}

func (i *indexer) markIncomplete(chainID ids.ID) error {
	key := make([]byte, hashing.HashLen+wrappers.ByteLen)
	copy(key, chainID[:])
	key[hashing.HashLen] = isIncompletePrefix
	return i.db.Put(key, nil)
}

// Returns true if this chain is incomplete
func (i *indexer) isIncomplete(chainID ids.ID) (bool, error) {
	key := make([]byte, hashing.HashLen+wrappers.ByteLen)
	copy(key, chainID[:])
	key[hashing.HashLen] = isIncompletePrefix
	return i.db.Has(key)
}

func (i *indexer) markPreviouslyIndexed(chainID ids.ID) error {
	key := make([]byte, hashing.HashLen+wrappers.ByteLen)
	copy(key, chainID[:])
	key[hashing.HashLen] = previouslyIndexedPrefix
	return i.db.Put(key, nil)
}

// Returns true if this chain is incomplete
func (i *indexer) previouslyIndexed(chainID ids.ID) (bool, error) {
	key := make([]byte, hashing.HashLen+wrappers.ByteLen)
	copy(key, chainID[:])
	key[hashing.HashLen] = previouslyIndexedPrefix
	return i.db.Has(key)
}

// Mark that the node has run at least once
func (i *indexer) markHasRun() error {
	return i.db.Put(hasRunKey, nil)
}

// Returns true if the node has run before
func (i *indexer) hasRun() (bool, error) {
	return i.db.Has(hasRunKey)
}
