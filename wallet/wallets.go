package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"os"
)

const walletFile = "./tmp/wallets.data"

type Wallets struct {
	//string would be our address.
	Wallets map[string]*Wallet
}

func CreateWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFile()
	return &wallets, err
}

func (ws *Wallets) AddWallet() string {
	wallet := MakeWallet()
	address := string(wallet.Address())
	ws.Wallets[address] = wallet
	return address
}

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) GetAllAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}
	return addresses
}

func (ws *Wallets) LoadFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	var wallets Wallets
	filecontent, err := ioutil.ReadFile(walletFile)

	Handle(err)

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(filecontent))
	err = decoder.Decode(&wallets)
	Handle(err)

	ws.Wallets = wallets.Wallets
	return nil
}

func (ws *Wallets) SaveFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	Handle(err)

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	Handle(err)
}
