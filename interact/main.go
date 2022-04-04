package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	todo "github.com/goABI/gen"
	"log"
	"math/big"
)

const (
	infuraRopstenTestNet = "https://ropsten.infura.io/v3/48284b00f92245f098e949dede474365"
	privateKey           = "私钥"
	contractAddr         = "0x1bab3A5577150E2beb6DBaF44edE5Fa8AA8489E5"
)

func main() {
	client, err := ethclient.Dial(infuraRopstenTestNet)
	if err != nil {
		fmt.Println("eth client: ", err)
		return
	}
	defer client.Close()

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	walletAddr1 := crypto.PubkeyToAddress(privKey.PublicKey)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("get chain id err: ", err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("get gas price err: ", err)
	}
	nonce, err := client.PendingNonceAt(context.Background(), walletAddr1)
	if err != nil {
		fmt.Println("get balance err: ", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		fmt.Println("new key err: ", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	cAddr := common.HexToAddress(contractAddr)
	t, err := todo.NewTodo(cAddr, client)
	if err != nil {
		fmt.Println("new todo err: ", err)
	}

	//add task
	//tx, err := t.Add(auth, "first task")
	//if err != nil {
	//	fmt.Println("add err: ", err)
	//}
	//fmt.Println(tx.Hash())

	//check add list
	//tasks, err := t.List(&bind.CallOpts{
	//	From: walletAddr1,
	//})
	//if err != nil {
	//	fmt.Println("list err: ", err)
	//}
	//fmt.Println(tasks)

	//update the task which id=0
	//tra, err := t.Update(auth, big.NewInt(0), "update task content")
	//if err != nil {
	//	fmt.Println("list err: ", err)
	//}
	//fmt.Println("update", tra.Hash())

	//toggle status
	//tra, err := t.Toggle(auth, big.NewInt(0))
	//if err != nil {
	//	fmt.Println("toggle err: ", err)
	//}
	//fmt.Println("Toggle tx", tra.Hash())

	//remove task
	tra, err := t.Remove(auth, big.NewInt(0))
	if err != nil {
		fmt.Println("remove task err: ", err)
	}
	fmt.Println("Remove task tx", tra.Hash())
}
