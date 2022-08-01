package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/aminmortezaie/golang-blockchain/blockchain"
)

//ecdsa stands for Elliptic Curve Digital Signature Algorithm
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	//see generate key func which is returning pointer of PrivateKey.
	private, err := ecdsa.GenerateKey(curve, rand.Reader)

	blockchain.Handle(err)

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}
