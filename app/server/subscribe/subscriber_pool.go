package subscribe

import (
  "fmt"
  // "encoding/json"
  // "time"
  	// "context"
  // 	"github.com/ethereum/go-ethereum/ethclient"
  // "github.com/ethereum/go-ethereum/core/types"
  // "github.com/go-redis/redis"
  "sync"
  	// "log"
)



type SubscriberPool struct {
    Subscribers []*Subscriber
    Blocks *BlockList
}

var subpool *SubscriberPool

func NewSubscriberPool() *SubscriberPool{
    blockList := NewBlockList()

    var subscribers []*Subscriber
    for _,host := range cfg.Networks {
        sb := NewSubscriber(host.Name,host.Http,host.WebSocket,blockList)
        subscribers = append(subscribers,sb)
    }

    subpool := &SubscriberPool{
      Subscribers:subscribers,
      Blocks: blockList,
    }
    return subpool
}

func (sp *SubscriberPool) Start(wg *sync.WaitGroup){
    for _,sub := range sp.Subscribers {
      fmt.Println("Start subscriber: ",sub.Name)
      sub.Start(wg)
    }
}
