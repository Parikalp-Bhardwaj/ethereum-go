package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Private Key", privateKey)

	// you can generate publickey by Elliptic Curve Digital Signature Algorithm
	getPublicKey := crypto.FromECDSA(privateKey)
	publicKey := hexutil.Encode(getPublicKey)
	fmt.Println("Public Key ", publicKey)
	//publicKey is a key now convert it into publicKey
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	fmt.Println("Public Address ", publicAddress)

}
