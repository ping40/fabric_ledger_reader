package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric/protoutil"

	"github.com/hyperledger/fabric/protos/common"
)

func showBlock(block *common.Block) {

	localBlock := &Block{
		Header: convertMyBlockHeader(block.Header),
		Data:   &BlockData{},
	}

	if true {
		dealBlockData(block, localBlock)

		localBlock.Metadata = showMetadata(block.Metadata)
	}

	blockJSON, _ := json.Marshal(localBlock)
	blockJSONString, _ := prettyprint(blockJSON)

	fmt.Println(string(blockJSONString))
	fmt.Println("\n\n  ...... \n\n")
}

func convertMyBlockHeader(header *common.BlockHeader) *BlockHeader {
	bh := &BlockHeader{
		Number:       header.Number,
		PreviousHash: header.PreviousHash,
		DataHash:     header.DataHash,
	}

	//	bh.MyHeaderHash = protoutil.BlockHeaderHash(header)

	return bh

}

func dealBlockData(block *common.Block, localBlock *Block) {
	for _, data := range block.Data.Data {
		if env, err := protoutil.GetEnvelopeFromBlock(data); err != nil {
			log.Fatal("error in GetEnvelopeFromBlock, %v", err)
		} else {
			if env == nil {
				localBlock.Data.Data = append(localBlock.Data.Data, &Envelope{
					Signature: "nil envelope",
				})

			} else {
				localBlock.Data.Data = append(localBlock.Data.Data, dealEnvelope(env))
			}
		}
	}
}
