package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) (block *Block) {
	block = &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
