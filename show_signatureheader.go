package main

import (
	"log"

	"github.com/hyperledger/fabric/protoutil"

	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/msp"
)

func showCreator(bytes []byte) *SerializedIdentity {

	sID := &msp.SerializedIdentity{}
	err := proto.Unmarshal(bytes, sID)
	if err != nil {
		log.Fatalf("Failed unmarshalling endorser: %v", err)
	}
	sID.IdBytes = showCert(sID.IdBytes)

	return &SerializedIdentity{
		Mspid:   sID.Mspid,
		IdBytes: sID.IdBytes,
	}
}

func showSignatureHeader(header *common.SignatureHeader) *SignatureHeader {
	s := &SignatureHeader{
		Nonce:   header.Nonce,
		Creator: showCreator(header.Creator),
	}

	return s
}

func getSignatureHeader(b []byte) *SignatureHeader {

	shdr, err := protoutil.GetSignatureHeader(b)
	if err != nil {
		log.Fatal("error, GetSignatureHeader :  ", err)
	}

	return showSignatureHeader(shdr)
}
