// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"math"

	"github.com/lasthyphen/dijetsnodego/codec"
	"github.com/lasthyphen/dijetsnodego/codec/linearcodec"
	"github.com/lasthyphen/dijetsnodego/utils/wrappers"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/txs"
)

var Codec codec.Manager

func init() {
	gc := linearcodec.NewCustomMaxLength(math.MaxInt32)
	Codec = codec.NewManager(math.MaxInt32)

	// To maintain codec type ordering, skip positions
	// for Proposal/Abort/Commit/Standard/Atomic blocks
	gc.SkipRegistrations(5)

	errs := wrappers.Errs{}
	errs.Add(
		txs.RegisterUnsignedTxsTypes(gc),
		Codec.RegisterCodec(txs.Version, gc),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}