// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nftfx

import (
	"testing"

	"github.com/lasthyphen/dijetsnodego/vms/components/verify"
)

func TestMintOutputState(t *testing.T) {
	intf := interface{}(&MintOutput{})
	if _, ok := intf.(verify.State); !ok {
		t.Fatalf("should be marked as state")
	}
}