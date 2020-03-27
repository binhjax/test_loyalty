package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "strings"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
		"github.com/binhnt-teko/test_loyalty/app/test/contracts"
)

func main() {
    client, err := ethclient.Dial("http://localhost:8502")
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress("0x0b1F50a633B951744379e5Da9c913DCD705cF6cD")
    query := ethereum.FilterQuery{
        FromBlock: big.NewInt(0),
        ToBlock:   big.NewInt(100000),
        Addresses: []common.Address{
            contractAddress,
        },
    }

    logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }

    contractAbi, err := abi.JSON(strings.NewReader(string(contracts.BusinessABI)))
    if err != nil {
        log.Fatal(err)
    }

    for _, vLog := range logs {

        fmt.Println("block = ", vLog.BlockHash.Hex(),vLog.BlockNumber,vLog.TxHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8

        event := new(contracts.BusinessEventCreateStash)
        err := contractAbi.Unpack(event, "event_createStash", vLog.Data)
        if err != nil {
            log.Fatal(err)
            fmt.Println("Error in event parse: ", err)
        }

        fmt.Println("Wallet: ", fmt.Sprintf("%s", event.WalletCode),event.WalletAddress.Hex(), event.WalletType) // foo

        var topics [4]string
        for i := range vLog.Topics {
            topics[i] = vLog.Topics[i].Hex()
        }

        fmt.Println("Topic:" , topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
    }

    eventSignature := []byte("ItemSet(bytes32,bytes32)")
    hash := crypto.Keccak256Hash(eventSignature)
    fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
