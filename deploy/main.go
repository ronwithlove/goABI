package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	todo "github.com/goABI/gen"
	"log"
	"math/big"
)

const (
	infuraRopstenTestNet = "https://api.s0.b.hmny.io"
	privateKey           = "25d00fb5bdd850498a3a654423d5e5b0d940d5d3849e32187671679950a68f12"
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

	addr, tx, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		fmt.Println("deploy contract err: ", err)
	}
	fmt.Println("----------------------------")
	fmt.Println(addr.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("----------------------------")

}
