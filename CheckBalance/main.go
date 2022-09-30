package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/977513988c264b3694ca2e751aa55ef1"
var gancheURL = "http://127.0.0.1:7545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatal("Error to create a ether client ", err)
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}
	// defer client.close()
	//BlockByNumber will give us latest block number
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("Error to create a ether client ", err)
	}

	fmt.Println(block.Number())

	add := "0x8f7b647493dF5259c435c14770d8a82F09cC43A8"
	address := common.HexToAddress(add)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal("Error to create a ether client ", err)
	}

	//Balance of address is
	fmt.Println("The Balance ", balance)

	// 1 ether = 10^18 wei convert wei to ether
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println("Ether is ", fBalance)
	//convert balance wei to ether
	balanceEther := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("Ether balance is ", balanceEther)
}
