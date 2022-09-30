package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"
	// newAccount, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(newAccount.Address)
	//Now How can we decrypt the file
	file, err := ioutil.ReadFile("./wallet/UTC--2022-08-23T14-22-39.899908916Z--9c80cc530f5b3acf80ae926d1171393cb0c1c389")

	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(file, password)
	if err != nil {
		log.Fatal(err)
	}
	//PrivateKey
	privateData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println(hexutil.Encode(privateData))

	//With this privateKey we are going to generate publicKey

	//PublicKey
	publicData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println(hexutil.Encode(publicData))

	//publicAddress
	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println("Address ", address)

}
