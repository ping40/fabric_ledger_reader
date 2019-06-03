package main

import (
	"log"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
)

func dealTxRwSet(bytes []byte) *rwsetutil.TxRwSet {

	rwset := &rwsetutil.TxRwSet{}
	if err := rwset.FromProtoBytes(bytes); err != nil {
		log.Fatalf("txRWSet.FromProtoBytes failed on %v", err)
	}

	return rwset
}
