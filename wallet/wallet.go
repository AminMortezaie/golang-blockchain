package wallet

import "crypto/ecdsa"

//ecdsa stands for Elliptic Curve Digital Signature Algorithm
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}
