package main

import (
	"fmt"
	"github.com/AntonVolov/cryptit/decrypt"
	"github.com/AntonVolov/cryptit/encrypt"
)

func main() {
	fmt.Println(encrypt.Nimbus("KodeKloud"))
	fmt.Println(decrypt.Bimbus("NrghNorxg"))
}


