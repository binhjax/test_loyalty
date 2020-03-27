package admin

import (
	// "os"
	"context"
	// "log"
	"github.com/ethereum/go-ethereum/rpc"
	// "github.com/ethereum/go-ethereum/p2p"
	"fmt"
	"strconv"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type RPCTransaction struct {
	BlockHash        common.Hash     `json:"blockHash"`
	BlockNumber      *hexutil.Big    `json:"blockNumber"`
	From             common.Address  `json:"from"`
	Gas              hexutil.Uint64  `json:"gas"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	Hash             common.Hash     `json:"hash"`
	Input            hexutil.Bytes   `json:"input"`
	Nonce            hexutil.Uint64  `json:"nonce"`
	To               *common.Address `json:"to"`
	TransactionIndex hexutil.Uint    `json:"transactionIndex"`
	Value            *hexutil.Big    `json:"value"`
	V                *hexutil.Big    `json:"v"`
	R                *hexutil.Big    `json:"r"`
	S                *hexutil.Big    `json:"s"`
}

func GetPool(server string) error {
	  ctx := context.Background()

		client, err := rpc.DialContext(ctx, server)
		if err != nil {
			fmt.Println("Unable to connect to network:%v\n", err)
			return err
		}
		method := "txpool_content"
		var result map[string]map[string]map[string][]*RPCTransaction
		client.CallContext(ctx, &result, method)

		min := uint64(0)
		//fmt.Println("Result:",result)
		for key, value := range result {
	    //fmt.Println("Key:", key)
			for key1, value1 := range value {
				//fmt.Println(key, key1)
				for key2, value2 := range value1 {
						nonce,_ := strconv.ParseUint(key2, 10, 64)
						if min == 0 {
							min = nonce
						}else{
							if nonce < min {
								 min = nonce
							}
						}
						fmt.Println(key, key1, key2, min, len(value2))
						for key3, value3 := range value2 {
							 fmt.Println(key, "-", key1,"-",key2,"-",key3,": Nonce: ",value3.Nonce)
						}
				}
			}
		}
		return nil
}
