package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const infuraRopstenTestNet = "https://ropsten.infura.io/v3/48284b00f92245f098e949dede474365"

func main() {
	client, err := ethclient.Dial(infuraRopstenTestNet)
	if err != nil {
		fmt.Println("eth client: ", err)
		return
	}
	defer client.Close()
	walletAddr := common.HexToAddress("0x398b02f35b6Cc21C43B40185228e1CEaDEac38CE")

	balance, err := client.BalanceAt(context.Background(), walletAddr, nil)
	if err != nil {
		fmt.Println("get balance err: ", err)
	}
	fmt.Println("The balance: ", balance)

}
