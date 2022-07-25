package main

import (
	"github.com/aminmortezaie/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("First block after Genesis.")
	chain.AddBlock("Second block after Genesis.")
	chain.AddBlock("Third block after Genesis.")

	for _, block := range chain.Blocks {
		block.DeriveHash()
		// fmt.Printf("PrevHash: %x\n", block.PrevHash)
		// fmt.Printf("block.Data: %s\n", block.Data)
	}
}
