package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const (
	infuraRopstenTestNet = "https://ropsten.infura.io/v3/48284b00f92245f098e949dede474365"
	privateKey           = "从metamask获取私钥"
)

func main() {
	client, err := ethclient.Dial(infuraRopstenTestNet)
	if err != nil {
		fmt.Println("eth client: ", err)
		return
	}
	defer client.Close()
	walletAddr1 := common.HexToAddress("0x398b02f35b6Cc21C43B40185228e1CEaDEac38CE")
	walletAddr2 := common.HexToAddress("0x1515EAF8971CE53495Ca41eDDb96E4155c904DAf")

	balance1, err := client.BalanceAt(context.Background(), walletAddr1, nil)
	if err != nil {
		fmt.Println("get balance err: ", err)
	}
	fmt.Println("The balance: ", balance1)

	balance2, err := client.BalanceAt(context.Background(), walletAddr2, nil)
	if err != nil {
		fmt.Println("get balance err: ", err)
	}
	fmt.Println("The balance: ", balance2)

	nonce, err := client.PendingNonceAt(context.Background(), walletAddr1)
	if err != nil {
		fmt.Println("get balance err: ", err)
	}
	amount := big.NewInt(1000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("get gas price err: ", err)
	}
	tx := types.NewTransaction(nonce, walletAddr2, amount, 21000, gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("get chain id err: ", err)
	}

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//publicKey := privKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("error casting public key to ECDSA")
	//}
	//
	//address := crypto.PubkeyToAddress(*publicKeyECDSA)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		fmt.Println("get sign tx err: ", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("get send tx err: ", err)
	}
	fmt.Println(signedTx.Hash().Hex())

}
