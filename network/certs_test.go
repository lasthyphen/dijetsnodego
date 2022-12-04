// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"crypto/tls"
	"sync"
	"testing"

	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/network/peer"
	"github.com/lasthyphen/dijetsnodego/staking"
)

var (
	certLock   sync.Mutex
	tlsCerts   []*tls.Certificate
	tlsConfigs []*tls.Config
)

func getTLS(t *testing.T, index int) (ids.NodeID, *tls.Certificate, *tls.Config) {
	certLock.Lock()
	defer certLock.Unlock()

	for len(tlsCerts) <= index {
		cert, err := staking.NewTLSCert()
		if err != nil {
			t.Fatal(err)
		}
		tlsConfig := peer.TLSConfig(*cert, nil)

		tlsCerts = append(tlsCerts, cert)
		tlsConfigs = append(tlsConfigs, tlsConfig)
	}

	cert := tlsCerts[index]
	return ids.NodeIDFromCert(cert.Leaf), cert, tlsConfigs[index]
}
