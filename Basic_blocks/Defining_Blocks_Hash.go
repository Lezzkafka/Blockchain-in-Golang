package main

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func hashfunction(block Block) string {
	header := string(block.Index) + block.Data + block.PrevHash + block.Timestamp
	hash_byte := sha256.Sum256([]byte(header))
	return hex.EncodeToString(hash_byte[:])
}

