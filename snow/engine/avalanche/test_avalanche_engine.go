// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"errors"

	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/snow/consensus/avalanche"
	"github.com/lasthyphen/dijetsnodego/snow/engine/common"
)

var (
	_ Engine = &EngineTest{}

	errGetVtx = errors.New("unexpectedly called GetVtx")
)

// EngineTest is a test engine
type EngineTest struct {
	common.EngineTest

	CantGetVtx bool
	GetVtxF    func(vtxID ids.ID) (avalanche.Vertex, error)
}

func (e *EngineTest) Default(cant bool) {
	e.EngineTest.Default(cant)
	e.CantGetVtx = false
}

func (e *EngineTest) GetVtx(vtxID ids.ID) (avalanche.Vertex, error) {
	if e.GetVtxF != nil {
		return e.GetVtxF(vtxID)
	}
	if e.CantGetVtx && e.T != nil {
		e.T.Fatalf("Unexpectedly called GetVtx")
	}
	return nil, errGetVtx
}