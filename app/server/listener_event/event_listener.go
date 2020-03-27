package eventlog

import (
    "context"
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/binhnt-teko/test_loyalty/app/server/config"
    "github.com/binhnt-teko/test_loyalty/app/server/contracts"
)

func Init() {
    if !contracts.Loaded {
       fmt.Println("--- Contract not loaded before. Please load")
       return
    }
    if contracts.LContract.Address == "" {
       fmt.Println("--- Contract is not deployed. Please deploy")
       return
    }
    cfg := config.Configuration
    endpoint := fmt.Sprintf("ws://%s/ws", cfg.Networks[0].WebSocket)
    client, err := ethclient.Dial(endpoint)
     if err != nil {
        fmt.Println("Cannot connect server ", err)
        // return "", err
         // log.Fatal(err)
     }


    contractAddress := common.HexToAddress(contracts.LContract.Address)
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }
    logs := make(chan types.Log)

    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("----- Start to listener logs from host ")
    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            fmt.Println(vLog) // pointer to event log
        }
    }
}
