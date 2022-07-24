package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
)

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) (pow *ProofOfWork) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow = &ProofOfWork{b, target}
	return
}

func (pow *ProofOfWork) InitData(nonce int) (data []byte) {
	data = bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return
}

//this would convert a num to HEX
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
