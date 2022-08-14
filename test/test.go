// this is just a test file.

package main

import (
	"fmt"

	"github.com/mr-tron/base58/base58"
)

func main() {
	f, _ := base58.Decode("1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq")

	fmt.Printf("%x", f)
}
