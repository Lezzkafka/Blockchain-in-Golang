package main

import "time"

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