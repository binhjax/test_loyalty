package admin

import (
		"fmt"
	  // "os"
		"strings"
     // "math/big"
		// "github.com/ethereum/go-ethereum/rpc"
		// "github.com/ethereum/go-ethereum/p2p"
	  "github.com/binhnt-teko/test_loyalty/app/server/config"
	  // "crypto/md5"
		// "encoding/hex"
		// "github.com/ethereum/go-ethereum/crypto"
		// "gopkg.in/yaml.v2"
	  // "io/ioutil"
		// wa "github.com/binhnt-teko/test_loyalty/app/server/account"
		"log"
		// "time"
	  // "bytes"
		"context"
    // "github.com/binhnt-teko/test_loyalty/app/server/connection"
    "github.com/binhnt-teko/test_loyalty/app/server/contracts"
    "github.com/ethereum/go-ethereum"
     "github.com/ethereum/go-ethereum/common"
       "github.com/ethereum/go-ethereum/accounts/abi"
     // "github.com/ethereum/go-ethereum/core/types"
     "github.com/ethereum/go-ethereum/ethclient"
)

func GetLog(FromBlock int64, ToBlock int64, Addresses []string ) ( string,error){
  cfg := config.Configuration
  endpoint := fmt.Sprintf("ws://%s/ws", cfg.Networks[0].WebSocket)
  client, err := ethclient.Dial(endpoint)
   if err != nil {
      fmt.Println("Cannot connect server ", err)
      return "", err
       // log.Fatal(err)
   }

   fmt.Println("GetLog: start ")
   addresses := []common.Address{}

   for _, addr := range Addresses {
     cAddress := common.HexToAddress(addr)
     addresses = append(addresses,cAddress)
   }
   // fmt.Println("Addresses: ",Addresses)

    query := ethereum.FilterQuery{
      // FromBlock: big.NewInt(FromBlock),
      // ToBlock:   big.NewInt(ToBlock),
      Addresses: addresses,
    }
    // conn := connection.Pool.GetConnection()

    logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
        fmt.Println("Cannot filter log from server", err)
        return "", err
    }
    fmt.Println("log:", logs)
    contractAbi, err := abi.JSON(strings.NewReader(string(contracts.ContractsABI)))
    if err != nil {
        fmt.Println("Cannot get smart contract ", err)
        return "", err
    }

    for _, vLog := range logs {
        fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
        fmt.Println(vLog.BlockNumber)     // 2394201
        fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

        event := new(contracts.ContractsEventRegisterAccETH)
        if err := contractAbi.Unpack(event, "event_registerAccETH", vLog.Data); err != nil {
          log.Fatal(err)
          return "", err
        }
        event.Raw = vLog
        fmt.Println(event.ListAcc)   // foo

        var topics [4]string
        for i := range vLog.Topics {
            topics[i] = vLog.Topics[i].Hex()
        }

        fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
      }
      return "",nil
}
