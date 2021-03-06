package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
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

func genblock(prevblock Block) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = prevblock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = prevblock.Hash
	newBlock.Hash = hashfunction(newBlock)

	return newBlock
}

func blocksvaildation(prevblock, newblock Block) bool {

	if prevblock.Index+1 != newblock.Index {
		return false
	}

	if prevblock.Hash != newblock.PrevHash {
		return false
	}

	if newblock.Hash != hashfunction(newblock) {
		return false
	}
	return true
}

func main() {
	t0 := time.Now().String()

	var b0 = Block{0, t0, "Let there be light", "_", "_"}
	fmt.Println(b0)

	time.Sleep(time.Second)

	b1 := Block(genblock(b0))
	fmt.Printf("Block %v, \n \t Time:%v; \n \t Hash: %v;\n", b1.Index, b1.Timestamp, b1.Hash)
	fmt.Printf("\t Check: %v.\n", blocksvaildation(b0, b1))

	time.Sleep(time.Second)

	b2 := Block(genblock(b1))
	fmt.Printf("Block %v, \n \t Time:%v; \n \t Hash: %v;\n", b2.Index, b2.Timestamp, b2.Hash)
	fmt.Printf("\t Check: %v.\n", blocksvaildation(b1, b2))

	time.Sleep(time.Second)

	b3 := Block(genblock(b2))
	fmt.Printf("Block %v, \n \t Time:%v; \n \t Hash: %v;\n", b3.Index, b3.Timestamp, b3.Hash)
	fmt.Printf("\t Check: %v.\n", blocksvaildation(b2, b3))
}
