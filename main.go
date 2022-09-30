package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://kovan.infura.io/v3/977513988c264b3694ca2e751aa55ef1"
)

func main() {

	client, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("Error to create a ether client ", err)
	}
	fmt.Println(block.Number())
	add := common.HexToAddress("0x19384d6dcAFb263C5fC5dF4E1DA03619c0102793")
	bal, err := client.BalanceAt(context.Background(), add, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bal)

}
