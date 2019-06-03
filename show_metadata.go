package main

import (
	"fmt"
	"log"

	"go.uber.org/zap/buffer"

	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger/fabric/protos/common"
	ab "github.com/hyperledger/fabric/protos/orderer"
	"github.com/hyperledger/fabric/protos/orderer/etcdraft"
)

func showMetadata(metadata *common.BlockMetadata) *BlockMetadata {
	bm := &BlockMetadata{
		//		Metadata:   metadata.Metadata,
		Metadata_1: showConfigMetadata(metadata.Metadata[common.BlockMetadataIndex_LAST_CONFIG]),
		Metadata_0: showSignatureMetadata(metadata.Metadata[common.BlockMetadataIndex_SIGNATURES]),
		Metadata_2: showTransactionsFilter(metadata.Metadata[common.BlockMetadataIndex_TRANSACTIONS_FILTER]),
		Metadata_3: showOrdererMetadata(metadata.Metadata[common.BlockMetadataIndex_ORDERER]),
	}

	return bm
}

func showOrdererMetadata(bytes []byte) *MetadataConsensus {
	mc := &MetadataConsensus{}

	raftMetadata := &etcdraft.BlockMetadata{}
	if err := proto.Unmarshal(bytes, raftMetadata); err != nil {
		kafkaMetadata := &ab.KafkaMetadata{}
		if err = proto.Unmarshal(bytes, kafkaMetadata); err != nil {
		} else {
			mc.ValueKafka = kafkaMetadata
		}
	} else {
		mc.ValueRaft = raftMetadata
	}
	return mc
}

func showTransactionsFilter(bytes []byte) string {
	var buf buffer.Buffer
	for _, b := range bytes {
		buf.AppendString(fmt.Sprintf("%x ", b))
	}
	return buf.String()
}

func showSignatureMetadata(bytes []byte) *Metadata {
	cm := &common.Metadata{}
	if err := proto.Unmarshal(bytes, cm); err != nil {
		log.Fatalf("error in Unmarshal, %v", err)
	}

	m := &Metadata{
		Signatures: dealSignature(cm.Signatures),
	}

	return m
}

func showConfigMetadata(bytes []byte) *Metadata {

	cm := &common.Metadata{}
	if err := proto.Unmarshal(bytes, cm); err != nil {
		log.Fatalf("error in Unmarshal, %v", err)
	}

	clc := &common.LastConfig{}
	if err := proto.Unmarshal(cm.Value, clc); err != nil {
		log.Fatalf("error in Unmarshal, %v", err)
	}

	m := &Metadata{
		Value:      clc,
		Signatures: dealSignature(cm.Signatures),
	}

	return m
}

func dealSignature(signatures []*common.MetadataSignature) []*MetadataSignature {
	var mss []*MetadataSignature
	if signatures == nil {
		return mss
	}

	for _, s := range signatures {
		sHdr := &common.SignatureHeader{}
		if err := proto.Unmarshal(s.SignatureHeader, sHdr); err != nil {
			log.Fatalf("error in Unmarshal, %v", err)
		}

		ms := &MetadataSignature{
			Signature:       s.Signature,
			SignatureHeader: showCreator(sHdr.Creator),
		}
		mss = append(mss, ms)

	}

	return mss
}
