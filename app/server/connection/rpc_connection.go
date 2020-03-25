
package connection

import (
    "github.com/ethereum/go-ethereum/ethclient"
    "sync"
    "fmt"
)

// var sha hash.Hash
type RpcConnection struct {
    Name string
	  Client  *ethclient.Client
  	Mux sync.Mutex
}

func NewRpcConnection(name string, url string) (*RpcConnection, error) {
    fmt.Println("Connect to host: ",url)
    cl, err  := ethclient.Dial("http://" + url)
    if err != nil {
       fmt.Println("Unable to connect to network: ", err)
       return nil, err
    }
    client := &RpcConnection{
      Name: name,
      Client: cl,
    }
    return client, nil
}
