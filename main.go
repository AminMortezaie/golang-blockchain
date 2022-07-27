package main

import (
	"fmt"
	"strconv"

	"github.com/aminmortezaie/golang-blockchain/blockchain"
)

type CommandLine struct {
	blockchain *blockchain.Blockchain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA -add a block to the chain")
	fmt.Println(" print -Prints the blocks in the chain")
}

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("First block after Genesis.")
	chain.AddBlock("Second block after Genesis.")
	chain.AddBlock("Third block after Genesis.")

	for _, block := range chain.Blocks {
		// block.DeriveHash()
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %v\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
