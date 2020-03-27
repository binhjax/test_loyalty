package main

import (
		"log"
		"github.com/ethereum/go-ethereum/ethclient"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		// "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
		// "github.com/ethereum/go-ethereum/rpc"
		"strings"
		"fmt"
		// "time"
		// "os"
		"math/big"
		"time"
		"context"
		"encoding/hex"
		"github.com/binhnt-teko/test_loyalty/app/test/contracts"
)

//const key  = `paste the contents of your JSON key file here`
const key  = `{"address":"ffbcd481c1330e180879b4d2b9b50642eea43c02","crypto":{"cipher":"aes-128-ctr","ciphertext":"351950aa30a37e4b385ae27ff2139c4151a6021333bd986602e80c2288f9e8fe","cipherparams":{"iv":"aec5c52378134e49a6037a5b77bec309"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8b0640866e9dbbba9f4a5da4348905b6f332a1b44a614ceadd0e9bd4ea7cdd7d"},"mac":"247c67172dbcdc48f031394ef1a25547f720b769be433a382a48028137f34002"},"id":"e52f52a5-cea4-459d-9e9b-ad8c76d7a562","version":3}`
func main(){
	// connect to an ethereum node  hosted by infura
	blockchain, err  := ethclient.Dial("http://localhost:8502")

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
	auth.GasLimit = 100000000

	//triggerAddr, _, trigger, err := DeployTrigger(auth, backends.NewRPCBackend(conn))
	// contractAddress, tx, _, err	:= contracts.DeployInbox(auth,blockchain)
	// contractAddress, tx, _, err	:= contracts.DeployMetaCoin(auth,blockchain,big.NewInt(1000))
	contractAddress, tx, _, err	:= contracts.DeployWalletBusiness(auth,blockchain)
	// contractAddress, tx, _, err	:= contracts.DeployTestCTX(auth,blockchain)
	if err != nil {
    log.Fatalf("Failed to deploy new trigger contract: %v", err)
  }
	fmt.Printf("Contract pending deploy: 0x%x\n", contractAddress)
	fmt.Printf("Transaction: 0x%x\n", tx.Hash())

	//Wait and get CodeAt address
	fmt.Println("Wait to get list of register accounts:.")
	time.Sleep(4*time.Second)

	bytecode, err := blockchain.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
			fmt.Println("Error in get bytecode: ", err)
	}
	if len(bytecode) > 0 {
			 fmt.Println("Data:",bytecode)
			 data := hex.EncodeToString(bytecode)
			 fmt.Println(data) // 60806...10029
	} else {
			 fmt.Println("is account, not smartcontract")
	}
}
