package tasks

import (
    "fmt"
    // "encoding/json"
    "strings"
    "time"
    "context"
    // 	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"
    // "github.com/go-redis/redis"
    // "sync"
    // "log"
    "github.com/binhnt-teko/test_loyalty/app/server/connection"
    "github.com/binhnt-teko/test_loyalty/app/server/workers"
    ttype "github.com/binhnt-teko/test_loyalty/app/server/types"
)

func ExecuteTransactions(header *types.Header) {
      conn := connection.Pool.GetConnection()
      block, err := conn.Client.BlockByHash(context.Background(), header.Hash())
      if err != nil {
        fmt.Println("Error block by hash: ",err)
        //Send to error channel
        return
      }
      // t := time.Now()
      // fmt.Println(t.Format(time.RFC822),"Block Number: ", header.Number.String(),"number of transactions:", len(block.Transactions()), " header hash: " , header.Hash().Hex())
      coinbase := block.Coinbase()
      for _, transaction := range block.Transactions(){
           nonce := transaction.Nonce()
           key := strings.TrimPrefix(transaction.Hash().Hex(),"0x")
           fmt.Println("Store data to queue: ", coinbase.Hex(),nonce,key )
           // workers.Pub.Publish()
           result := ttype.TransactionResult{
              Nonce: nonce,
              TransactionHash: key,
           }
           err = workers.Pub.Publish("result",result)
           if err != nil {
              fmt.Println(time.Now(),"send error to error queue", err)
           }

      }
}
