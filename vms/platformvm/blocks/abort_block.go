// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocks

import (
	"time"

	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/snow"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/txs"
)

var (
	_ BlueberryBlock = &BlueberryAbortBlock{}
	_ Block          = &ApricotAbortBlock{}
)

type BlueberryAbortBlock struct {
	Time              uint64 `serialize:"true" json:"time"`
	ApricotAbortBlock `serialize:"true"`
}

func (b *BlueberryAbortBlock) Timestamp() time.Time  { return time.Unix(int64(b.Time), 0) }
func (b *BlueberryAbortBlock) Visit(v Visitor) error { return v.BlueberryAbortBlock(b) }

func NewBlueberryAbortBlock(
	timestamp time.Time,
	parentID ids.ID,
	height uint64,
) (*BlueberryAbortBlock, error) {
	blk := &BlueberryAbortBlock{
		Time: uint64(timestamp.Unix()),
		ApricotAbortBlock: ApricotAbortBlock{
			CommonBlock: CommonBlock{
				PrntID: parentID,
				Hght:   height,
			},
		},
	}
	return blk, initialize(blk)
}

type ApricotAbortBlock struct {
	CommonBlock `serialize:"true"`
}

func (b *ApricotAbortBlock) initialize(bytes []byte) error {
	b.CommonBlock.initialize(bytes)
	return nil
}

func (*ApricotAbortBlock) InitCtx(ctx *snow.Context) {}

func (*ApricotAbortBlock) Txs() []*txs.Tx          { return nil }
func (b *ApricotAbortBlock) Visit(v Visitor) error { return v.ApricotAbortBlock(b) }

func NewApricotAbortBlock(
	parentID ids.ID,
	height uint64,
) (*ApricotAbortBlock, error) {
	blk := &ApricotAbortBlock{
		CommonBlock: CommonBlock{
			PrntID: parentID,
			Hght:   height,
		},
	}
	return blk, initialize(blk)
}