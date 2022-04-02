package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
)

const BscTestNet = "ddd"

func callContractWithAbi(client *ethclient.Client, privKey *ecdsa.PrivateKey, from, to common.Address, contract string) (string, error) {
	//create tx
	nonce, err := client.NonceAt(context.Background(), from, nil)
	if err != nil {
		fmt.Println("get nonce: ", err)
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("get price: ", err)
		return "", err
	}

	//function data
	abiData, err := ioutil.ReadFile("abi.txt")
	if err != nil {
		fmt.Println("get abi: ", err)
		return "", err
	}
	contractABI, err := abi.JSON(bytes.NewReader(abiData))
	if err != nil {
		fmt.Println("read abi ", err)
		return "", err
	}
	amount, _ := new(big.Int).SetString("1000000000000000000", 10)
	callData, err := contractABI.Pack("transfer", to, amount)
	if err != nil {
		fmt.Println("abi pack: ", err)
		return "", err
	}

	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(contract),
		big.NewInt(0),
		uint64(30000),
		gasPrice,
		callData,
	)
	//sign tx
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(97)), privKey)
	if err != nil {
		fmt.Println("sign tx: ", err)
		return "", err
	}

	//sent tx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("sent tx: ", err)
		return "", err
	}

	return signedTx.Hash().Hex(), nil

}

func main() {
	client, err := ethclient.Dial(BscTestNet)
}
