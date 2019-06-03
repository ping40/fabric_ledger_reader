package main

import (
	"log"

	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protoutil"
)

func dealEnvelope(envelope *common.Envelope) *Envelope {
	e := &Envelope{}

	payload, err := protoutil.GetPayload(envelope)
	if err != nil {
		log.Fatal("error, protoutil.GetPayload :  ", err)
	}
	p := &Payload{
		//Data:   payload.Data,
		Header: &Header{},
	}

	e.Payload = p

	chdr, err := protoutil.UnmarshalChannelHeader(payload.Header.ChannelHeader)
	if err != nil {
		log.Fatal("error, UnmarshalChannelHeader :  ", err)
	}

	p.Header.ChannelHeader = chdr

	p.Header.SignatureHeader = getSignatureHeader(payload.Header.SignatureHeader)

	switch common.HeaderType(chdr.Type) {
	case common.HeaderType_ENDORSER_TRANSACTION:

		p.Endorser_transaction = dealTransaction(payload.Data)

	case common.HeaderType_ORDERER_TRANSACTION:
		// TODO
	case common.HeaderType_CONFIG:

		p.Config_envelope = dealConfigEnvelope(payload)

	default:
		log.Fatalf("Illegail block header type:, %v", chdr.Type)
	}

	return e
}
