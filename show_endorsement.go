package main

import (
	"github.com/hyperledger/fabric/protos/peer"
)

func showEndorsements(endorsements []*peer.Endorsement) []*SerializedIdentity {
	var sIDs []*SerializedIdentity

	for _, e := range endorsements {
		sIDs = append(sIDs, showCreator(e.Endorser))
	}

	return sIDs
}
