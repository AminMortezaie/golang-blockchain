package wallet

import (
	"github.com/aminmortezaie/golang-blockchain/blockchain"
	"github.com/mr-tron/base58"
)

func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)
	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	blockchain.Handle(err)
	return decode
}
