package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//Private Key
	keyGenerate, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//Generate Private Key With from ECDSA
	privateData := crypto.FromECDSA(keyGenerate)
	//let encode data
	privateEncode := hexutil.Encode(privateData)
	fmt.Println("Private Address ", privateEncode)
	//now create PublicKey
	// you can generate publickey by Elliptic Curve Digital Signature Algorithm (ECDSA)
	publicData := crypto.FromECDSAPub(&keyGenerate.PublicKey)
	publicEncode := hexutil.Encode(publicData)
	fmt.Println("Public Key ", publicEncode)
	// generate the user’s public address from the public key
	publicAddress := crypto.PubkeyToAddress(keyGenerate.PublicKey).Hex()
	fmt.Println("public Address ", publicAddress)
}
