package main

import (
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"
)

func nextBlockBytes() ([]byte, error) {
	var lenBytes []byte
	var err error

	// At the end of file
	if fileOffset == fileSize {
		return nil, nil
	}

	remainingBytes := fileSize - fileOffset
	peekBytes := 8
	if remainingBytes < int64(peekBytes) {
		peekBytes = int(remainingBytes)
	}
	if lenBytes, err = fileReader.Peek(peekBytes); err != nil {
		return nil, err
	}

	length, n := proto.DecodeVarint(lenBytes)
	if n == 0 {
		return nil, fmt.Errorf("Error in decoding varint bytes [%#v]", lenBytes)
	}

	bytesExpected := int64(n) + int64(length)
	if bytesExpected > remainingBytes {
		return nil, fmt.Errorf("corrupted file")
	}

	// skip the bytes representing the block size
	if _, err = fileReader.Discard(n); err != nil {
		return nil, err
	}

	blockBytes := make([]byte, length)
	if _, err = io.ReadAtLeast(fileReader, blockBytes, int(length)); err != nil {
		return nil, err
	}

	fileOffset += int64(n) + int64(length)
	return blockBytes, nil
}
