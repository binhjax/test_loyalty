package listeners

import (
    "fmt"
    "context"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"
    // "sync"
    "log"
)
type BlockListener struct {
   Name string
   SocketUrl string
   Active bool
}

func NewBlockListener(name string, httpUrl string,socketUrl string,)  *BlockListener {
     listener :=  &BlockListener{
       Name: name,
       SocketUrl: socketUrl,
       Active: true,
     }
     return listener
}
func (sb *BlockListener) ListenEvent(){
		fmt.Println("Subscriber:", sb.Name ,"Listening from: ", sb.SocketUrl)
		websocket, err := ethclient.Dial("ws://" + sb.SocketUrl)
		if err != nil {
				fmt.Println("Cannot connect to websocket: ", err)
				return
		}
		headers := make(chan *types.Header)
		sub, err := websocket.SubscribeNewHead(context.Background(), headers)
		if err != nil {
		    fmt.Println("Cannot SubscribeNewHead to host: ", sb.SocketUrl ," Error: ",err)
				return
		}
	  fmt.Println("Start listening: ",sb.SocketUrl,"  ")
		for {
					select {
								case err := <-sub.Err():
										fmt.Println("Error from: ",sb.SocketUrl," Error: ",err)
										log.Fatal(err)
								case header := <-headers:
                   // fmt.Println("1. Block Number: ", header.Number.String()," Subscriber: ", sb.Name, " -> call CheckHeader")
                    // Process header
                    go func(){
                        CheckFinalize(header,sb.Name)
                    }()
						}
		}
}
