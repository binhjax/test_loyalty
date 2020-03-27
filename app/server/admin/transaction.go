package admin

import (
	// "os"
	"context"
	"fmt"
	// "log"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"
		"github.com/binhnt-teko/test_loyalty/app/server/connection"
		  "github.com/ethereum/go-ethereum/core/types"
)

func GetTransactionState(transHash string ) (*types.Receipt, error) {
	// client, err  := ethclient.Dial(webserver)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	conn := connection.Pool.GetConnection()

	tx_hash := common.HexToHash("0x"+transHash)
	tx, isPending, err := conn.Client.TransactionByHash(context.Background(), tx_hash)
	if err != nil {
	 	return nil, err
	}

	fmt.Println("Transaction status: pending = ",isPending,", hash: ", tx.Hash()) // 1
	receipt, err := conn.Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
			return nil, err
	}
	fmt.Println("GasUsed:",receipt.GasUsed) // 1
	return receipt, nil
}
