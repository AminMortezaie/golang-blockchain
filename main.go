package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Blockchain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) (block *Block) {
	block = &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockchain()

	chain.AddBlock("First block after Genesis.")
	chain.AddBlock("Second block after Genesis.")
	chain.AddBlock("Third block after Genesis.")

	for _, block := range chain.blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("block.Data: %s\n", block.Data)
	}
}
