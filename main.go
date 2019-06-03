package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	file       *os.File
	fileSize   int64
	fileOffset int64
	fileReader *bufio.Reader

	fileName   = flag.String("file_name", "", "The qualified file name of ledger file, for example : /tmp/blockfile_000000")
	numberFrom = flag.Uint64("number_from", 0, "The first block number to print")
	numberEnd  = flag.Uint64("number_end", 0, "The last block number to print")
)

func checkParameters() {
	if *fileName == "" {
		log.Fatalf("Please input filename")
	}
	if *numberFrom < 0 {
		*numberFrom = 0
	}
	if *numberEnd < *numberFrom {
		*numberEnd = *numberFrom
	}
}

func main() {

	flag.Parse()

	checkParameters()

	var err error
	if file, err = os.OpenFile(*fileName, os.O_RDONLY, 0600); err != nil {
		fmt.Printf("error in OpenFile,[%s], error=[%v]\n", fileName, err)
		return
	}
	defer file.Close()

	if fileInfo, err := file.Stat(); err != nil {
		fmt.Printf("error in fileStat, [%s], error=[%v]\n", fileName, err)
		return
	} else {
		fileOffset = 0
		fileSize = fileInfo.Size()
		fileReader = bufio.NewReader(file)
	}

	for {
		if blockBytes, err := nextBlockBytes(); err != nil {
			fmt.Printf("error in  nextBlockBytes, [%s], error=[%v]\n", fileName, err)
			break
		} else if blockBytes == nil {
			// End of file
			break
		} else {
			if block, err := deserializeBlock(blockBytes); err != nil {
				fmt.Printf("ERROR: Cannot deserialize block from file: [%s], error=[%v]\n", fileName, err)
				break
			} else {
				if block.Header.Number >= *numberFrom && block.Header.Number <= *numberEnd {
					showBlock(block)
				}
			}
		}
	}

}
