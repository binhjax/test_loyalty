package main

import (
	// "os"
	"log"
	"github.com/binhnt-teko/test_loyalty/app/test/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"strings"
	// "bytes"
	// "io/ioutil"
	"fmt"
	"math/big"
	"context"
	"time"
)

//const key  = `paste the contents of your JSON key file here`
const key  = `{"address":"ffbcd481c1330e180879b4d2b9b50642eea43c02","crypto":{"cipher":"aes-128-ctr","ciphertext":"351950aa30a37e4b385ae27ff2139c4151a6021333bd986602e80c2288f9e8fe","cipherparams":{"iv":"aec5c52378134e49a6037a5b77bec309"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8b0640866e9dbbba9f4a5da4348905b6f332a1b44a614ceadd0e9bd4ea7cdd7d"},"mac":"247c67172dbcdc48f031394ef1a25547f720b769be433a382a48028137f34002"},"id":"e52f52a5-cea4-459d-9e9b-ad8c76d7a562","version":3}`

func main(){
	contractAddr := "0xd04abbd25fd70b1fc97a9f59c895a3494d273150"
	client, err  := ethclient.Dial("http://localhost:8502")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	gasPrice := new(big.Int)
	gasPrice.SetString("1",10)
	auth.GasPrice = gasPrice
	auth.GasLimit = 1000000

	contractAddress := common.HexToAddress(contractAddr)
	instance, err := contracts.NewTestCTX(contractAddress, client)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	val := int64(-10)
	fmt.Println("Value: ", val)
	tx, err := instance.Test(auth, big.NewInt(val))
	if err != nil {
			log.Fatalf(" Transaction create error: ", err)
	}
	fmt.Println(" Transaction =",tx.Hash().Hex())

	fmt.Println("Wait to get tramsaction: ")
	time.Sleep(1000*time.Millisecond)

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		fmt.Println(" Error get receipt: ",err)
		return
	}
	fmt.Println("GasUsed:",receipt.GasUsed) // 1
	fmt.Println("Status:",receipt.Status) // 1
	fmt.Println("Status:",receipt.TxHash.Hex()) // 1
}
