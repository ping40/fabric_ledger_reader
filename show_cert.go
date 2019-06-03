package main

import "crypto/sha256"

func showCert(cert []byte) []byte {
	h := sha256.New()
	h.Write(cert)
	return h.Sum(nil)[0:16]
}
