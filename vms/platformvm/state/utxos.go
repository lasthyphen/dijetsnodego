// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/vms/components/djtx"
)

type UTXOGetter interface {
	GetUTXO(utxoID ids.ID) (*djtx.UTXO, error)
}

type UTXOAdder interface {
	AddUTXO(utxo *djtx.UTXO)
}

type UTXODeleter interface {
	DeleteUTXO(utxoID ids.ID)
}
