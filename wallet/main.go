package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	privakey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privakeyData := crypto.FromECDSA(privakey)
	fmt.Println(hexutil.Encode(privakeyData))

	pubicKeyData := crypto.FromECDSAPub(&privakey.PublicKey)
	fmt.Println(hexutil.Encode(pubicKeyData))

	fmt.Println(crypto.PubkeyToAddress(privakey.PublicKey).Hex())

}
