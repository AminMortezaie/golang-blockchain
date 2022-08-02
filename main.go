package main

import (
	"os"

	"github.com/aminmortezaie/golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()

	// w := wallet.MakeWallet()
	// w.Address()
}
