package main

import (
	"fmt"

	"github.com/igweo/blockchain-go/blockchain"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func main() {
	chain := blockchain.InitalizeBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
