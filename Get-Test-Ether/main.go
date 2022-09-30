package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://kovan.infura.io/v3/977513988c264b3694ca2e751aa55ef1"

func main() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// password := "password"
	// newAccount, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(newAccount.Address)

	client, err := ethclient.DialContext(context.Background(), infuraURL)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// address1 = 9c80cc530f5b3acf80ae926d1171393cb0c1c389
	// address2 = c736ccad380e922693a790f36c59f644ebd3420c

	address1 := common.HexToAddress("0x19384d6dcAFb263C5fC5dF4E1DA03619c0102793")
	address2 := common.HexToAddress("c736ccad380e922693a790f36c59f644ebd3420c")

	balance1, err := client.BalanceAt(context.Background(), address1, nil)

	if err != nil {
		log.Fatal(err)
	}

	balance2, err := client.BalanceAt(context.Background(), address2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance 1", balance1)
	fmt.Println("Balance 2", balance2)

	nonce, err := client.PendingNonceAt(context.Background(), address1)
	if err != nil {
		log.Fatal(err)
	}
	amount := big.NewInt(1000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, address2, amount, 21000, gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("./wallet/UTC--2022-08-24T08-56-48.814296491Z--c736ccad380e922693a790f36c59f644ebd3420c")
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx send: %s", tx.Hash().Hex())
}
