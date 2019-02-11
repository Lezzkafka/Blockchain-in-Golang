package main

import (
	"fmt"
	"time"
)

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