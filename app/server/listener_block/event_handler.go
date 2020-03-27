package listeners

import (
    // "fmt"
    // "encoding/json"
    // "strings"
    // "time"
    // "context"
    // 	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"
    // "github.com/go-redis/redis"
    // "sync"
    // "log"
    // "github.com/binhnt-teko/test_loyalty/app/server/connection"
    "github.com/binhnt-teko/test_loyalty/app/server/tasks"
)
 func LoadTransactions(header *types.Header) error {
     // blNumber := header.Number.String()
     // fmt.Println("4. Try to load transaction from block: " + blNumber)
     tasks.ExecuteTransactions(header)
     return nil
 }
