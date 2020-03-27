package connection

import (
   // "context"
   "time"
   "fmt"
   "sync"
   // "errors"
   // "strings"
    // "github.com/ethereum/go-ethereum/core/types"
    // "github.com/ethereum/go-ethereum/common"
    // "github.com/ethereum/go-ethereum"
    // "math/big"
    // "math"
    "github.com/binhnt-teko/test_loyalty/app/server/config"
)

type RpcPool struct {
  Nodes []*EthNode
  Current int
  mutex sync.Mutex
  Mode int
  AccountNode map[string]*EthNode
}

var Pool *RpcPool

func Init() {
    cfg := config.Configuration
    max_conns := cfg.Webserver.MaxRpcConnection
    mode := cfg.Webserver.RoutingMode
    hosts := cfg.Networks

    var nodes []*EthNode
    for _,host := range hosts {
       node := NewEthNode(host.Name, max_conns, host.Http)
       nodes = append(nodes,node)
     }
     accountNode := make(map[string]*EthNode)

     Pool =  &RpcPool{
        Nodes: nodes,
        Current: -1,
        Mode: mode,
        AccountNode: accountNode,
     }
}

func (r *RpcPool) Process(){
      for {
          select {
              case <-time.After(5*time.Second):
                   go func() {
                     r.UpdateHealth()
                   }()
               }
        }
}
func (r *RpcPool) UpdateHealth(){
    r.mutex.Lock()
    defer r.mutex.Unlock()
    fmt.Println("Process to check health of node. To automatically active node")
    for _, node := range r.Nodes {
      if !node.Active {
          node.UpdateHealth()
      }
   }
}
func (r *RpcPool) DeactiveNode(name string)  {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    for _, node := range r.Nodes {
      if node.Name == name {
         node.Active  = false
      }
    }
}
func (r *RpcPool) ActiveNodeLength() int {
   n := 0
   for _, node := range r.Nodes {
      if node.Active {
         n = n + 1
      }
   }
   return n
}
func (r *RpcPool) GetConnectionByAccount(addr string)  (*RpcConnection) {
    r.mutex.Lock()
    defer r.mutex.Unlock()
    nNode := r.ActiveNodeLength()
    if nNode >0 {
        node, ok := r.AccountNode[addr]
        if !ok  || node.Active == false {
            fmt.Println("RpcPool.GetConnectionByAccount: Update map: Add account: ", addr)
            nAccounts := len(r.AccountNode)
            idx := nAccounts % nNode
            node = r.Nodes[idx]
            r.AccountNode[addr] = node
        }
        return node.GetConnection()
    }
    return nil
}
func (r *RpcPool) GetConnection() (*RpcConnection) {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    len := len(r.Nodes)
    retry := 0
    if len >0 {
        for retry < len {
          r.Current = (r.Current + 1 ) % len
          node := r.Nodes[r.Current]
          if node.Active {
            // fmt.Println("Get connection: ", r.Current, " Node: ", node.HttpUrl)
             return node.GetConnection()
          }
          retry = retry + 1
        }
    } else {
      fmt.Println("Not find node in list")
    }
    return nil
}

func (r *RpcPool) GetConnectionFix() (*RpcConnection) {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    len := len(r.Nodes)
    if len >0 {
      fmt.Println("Len : ", len, ", current: ", r.Current)
      if r.Current < 0 || r.Current >= len  {
        fmt.Println("Current < 0. Find first active node")
        for i, node := range r.Nodes {
          if node.Active {
             r.Current = i
             return node.GetConnection()
          }
        }
        return nil
      } else {
        node := r.Nodes[r.Current]
        if node.Active {
           return node.GetConnection()
        } else {
            fmt.Println("Current node not active. Find first active node")
            for i, node := range  r.Nodes {
              if node.Active {
                 r.Current = i
                 return node.GetConnection()
              }
            }
            return nil
        }
      }
    } else {
      fmt.Println("Not find node in list")
    }
    return nil
}
