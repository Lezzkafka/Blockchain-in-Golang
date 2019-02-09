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
	PrevHash  string
	Hash      string
}

func hashfunction(block Block) string {

	Input := string(block.Index) + block.PrevHash + block.Timestamp
	h := sha256.New()
	h.Write([]byte(Input))
	Output := h.Sum(nil)
	return hex.EncodeToString(Output)
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

	var b0 = Block{0, t0, "_", "_"}
	fmt.Println(b0)

	time.Sleep(time.Second)

	b1 := Block(genblock(b0))
	fmt.Println(b1)
	fmt.Println(blocksvaildation(b0, b1))

	time.Sleep(time.Second)

	b2 := Block(genblock(b1))
	fmt.Println(b2)
	fmt.Println(blocksvaildation(b1, b2))

	time.Sleep(time.Second)

	b3 := Block(genblock(b2))
	fmt.Println(b3)
	fmt.Println(blocksvaildation(b2, b3))
}
