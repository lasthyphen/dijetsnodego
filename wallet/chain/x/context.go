// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package x

import (
	stdcontext "context"

	"github.com/lasthyphen/dijetsnodego/api/info"
	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/vms/avm"
)

var _ Context = (*context)(nil)

type Context interface {
	NetworkID() uint32
	BlockchainID() ids.ID
	DJTXAssetID() ids.ID
	BaseTxFee() uint64
	CreateAssetTxFee() uint64
}

type context struct {
	networkID        uint32
	blockchainID     ids.ID
	djtxAssetID      ids.ID
	baseTxFee        uint64
	createAssetTxFee uint64
}

func NewContextFromURI(ctx stdcontext.Context, uri string) (Context, error) {
	infoClient := info.NewClient(uri)
	xChainClient := avm.NewClient(uri, "X")
	return NewContextFromClients(ctx, infoClient, xChainClient)
}

func NewContextFromClients(
	ctx stdcontext.Context,
	infoClient info.Client,
	xChainClient avm.Client,
) (Context, error) {
	networkID, err := infoClient.GetNetworkID(ctx)
	if err != nil {
		return nil, err
	}

	chainID, err := infoClient.GetBlockchainID(ctx, "X")
	if err != nil {
		return nil, err
	}

	asset, err := xChainClient.GetAssetDescription(ctx, "DJTX")
	if err != nil {
		return nil, err
	}

	txFees, err := infoClient.GetTxFee(ctx)
	if err != nil {
		return nil, err
	}

	return NewContext(
		networkID,
		chainID,
		asset.AssetID,
		uint64(txFees.TxFee),
		uint64(txFees.CreateAssetTxFee),
	), nil
}

func NewContext(
	networkID uint32,
	blockchainID ids.ID,
	djtxAssetID ids.ID,
	baseTxFee uint64,
	createAssetTxFee uint64,
) Context {
	return &context{
		networkID:        networkID,
		blockchainID:     blockchainID,
		djtxAssetID:      djtxAssetID,
		baseTxFee:        baseTxFee,
		createAssetTxFee: createAssetTxFee,
	}
}

func (c *context) NetworkID() uint32 {
	return c.networkID
}

func (c *context) BlockchainID() ids.ID {
	return c.blockchainID
}

func (c *context) DJTXAssetID() ids.ID {
	return c.djtxAssetID
}

func (c *context) BaseTxFee() uint64 {
	return c.baseTxFee
}

func (c *context) CreateAssetTxFee() uint64 {
	return c.createAssetTxFee
}
