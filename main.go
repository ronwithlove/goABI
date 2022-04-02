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
	"math"
	"math/big"
)

const (
	infuraRopstenTestNet = "https://ropsten.infura.io/v3/48284b00f92245f098e949dede474365"
	ContractAdd          = "0x944E2E2c632C4D6aF195DC3Bdec9C17F6fc6F600"
	ganacheURL           = "HTTP://127.0.0.1:7545"
	infuraMainNet        = "https://mainnet.infura.io/v3/48284b00f92245f098e949dede474365"
)

var PrivateKey1 *ecdsa.PrivateKey
var FromAddr common.Address
var ToAddr common.Address

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
	client, err := ethclient.Dial(infuraMainNet)
	if err != nil {
		fmt.Println("eth client: ", err)
		return
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println("get a block err: ", err)
	}
	fmt.Println(block.Number())

	address := common.HexToAddress(ContractAdd)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		fmt.Println("get balance err: ", err)
	}
	fmt.Println("The balance: ", balance)

	fBlance := new(big.Float)
	fBlance.SetString(balance.String())//科学计数法
	value:=new(big.Float).Quo(fBlance,big.NewFloat(math.Pow10(18)))
	fmt.Println("The balance: ", value)


	//txhash, err := callContractWithAbi(client, PrivateKey1, FromAddr, ToAddr, ContractAdd)
	//if err != nil {
	//	fmt.Println("call contract: ", err)
	//	return
	//}
	//fmt.Println(txhash)
}
