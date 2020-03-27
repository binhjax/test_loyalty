package listeners

import (
    "fmt"
    // "encoding/json"
    // "time"
    // "context"
    // 	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"
    // "github.com/go-redis/redis"
    "sync"
    // "log"
    "github.com/binhnt-teko/test_loyalty/app/server/config"
)
type eventHandler = func(header *types.Header) error

var Listeners []*BlockListener
var Blocks *BlockList
var EventHandlerList []eventHandler

func Init(){
    cfg := config.Configuration
    Blocks = NewBlockList()

    Listeners = make([]*BlockListener,0)
    for _,host := range cfg.Networks {
         // fmt.Println("Create listener to host: " + host.Name)
        listener := NewBlockListener(host.Name,host.Http,host.WebSocket)
        Listeners = append(Listeners,listener)
    }
    SubcribeEvenHandler(LoadTransactions)
}

func Start(wg *sync.WaitGroup){
  wg.Add(len(Listeners))
  for _,listener := range Listeners {
    // fmt.Println("Start listen block from : ",listener.Name)
    go func (p *BlockListener){
        fmt.Println("Start listen block from : ",p.Name)
        defer wg.Done()
        p.ListenEvent()
    }(listener)
  }
}

func CheckFinalize(header *types.Header, Name string){
    //Query redis
    blNumber := header.Number.String()

    // fmt.Println("2. BlockListener:",Name,"Check block: ",blNumber)
    if value, ok := Blocks.Get(blNumber); ok {
         if value != Name {
             // fmt.Println(" 3. Call worker to get transaction from block:",blNumber)
             if len(EventHandlerList) > 0 {
               for _,handler := range EventHandlerList {
                  handler(header)
               }
             }
         }else{
           // fmt.Println("3. Same BlockListener received same block :",blNumber)
         }
    } else {
        // fmt.Println("3. Not find blockNumber:",blNumber)
        Blocks.Set(blNumber,Name)
    }
}
func SubcribeEvenHandler(handler eventHandler){
    EventHandlerList = append(EventHandlerList,handler)
}
