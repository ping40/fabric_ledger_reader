package main

import (
	"log"

	"github.com/hyperledger/fabric/protos/peer"

	"github.com/hyperledger/fabric/protoutil"
)

func dealTransaction(data []byte) *Transaction {
	t := &Transaction{}

	tx, err := protoutil.GetTransaction(data)
	if err != nil {
		log.Fatal("error GetTransaction, %v", err)
	}
	for _, act := range tx.Actions {
		ta := &TransactionAction{}

		ta.Header = getSignatureHeader(act.Header)

		ccActionPayload, err := protoutil.GetChaincodeActionPayload(act.Payload)
		if err != nil {
			log.Fatal("error GetChaincodeActionPayload, %v", err)
		}

		ta.Payload = &ChaincodeActionPayload{
			ChaincodeProposalPayload: dealChaincodeProposalPayload(ccActionPayload.ChaincodeProposalPayload),
			Action: dealAction(ccActionPayload.Action),
		}

		t.Actions = append(t.Actions, ta)
	}

	return t
}

func dealAction(action *peer.ChaincodeEndorsedAction) *ChaincodeEndorsedAction {
	cea := &ChaincodeEndorsedAction{
		ProposalResponsePayload: dealChaincodeProposalPayload(action.ProposalResponsePayload),
		Endorsements:            showEndorsements(action.Endorsements),
	}

	return cea
}

func dealChaincodeProposalPayload(bytes []byte) *ProposalResponsePayload {
	// extract the proposal response payload
	prp, err := protoutil.GetProposalResponsePayload(bytes)
	if err != nil {
		log.Fatal("error GetProposalResponsePayload, %v", err)
	}
	return &ProposalResponsePayload{
		ProposalHash: prp.ProposalHash,
		Extension:    dealChaincodeAction(prp.Extension),
	}

}

func dealChaincodeAction(bytes []byte) *ChaincodeAction {
	ca := &ChaincodeAction{}
	if bytes == nil {
		return ca
	}

	respPayload, err := protoutil.GetChaincodeAction(bytes)
	if err != nil {
		log.Fatalf("GetChaincodeAction error %v", err)
	}
	ca.ChaincodeId = respPayload.ChaincodeId
	ca.Events = respPayload.Events
	ca.Response = respPayload.Response
	ca.Results = dealTxRwSet(respPayload.Results)
	ca.TokenOperations = respPayload.TokenOperations

	return ca

}
