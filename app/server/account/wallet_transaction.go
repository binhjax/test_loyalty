
package account

import (
   "context"
   "time"
   "fmt"
   // "sync"
   "errors"
   "strings"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum"
    "math/big"
    "math"
    "github.com/binhnt-teko/test_loyalty/app/server/connection"
)

func (r *WalletAccount) SendTransaction(signedTx *types.Transaction,  nonce uint64) (error){
      retry := 0
      for retry < 3 {
          conn := connection.Pool.GetConnection()
          if conn == nil {
             return errors.New("Not find connection. Please check")
          }
          //Lock rpc connection before send
          conn.Mux.Lock()
          defer 	conn.Mux.Unlock()
          err := conn.Client.SendTransaction(context.Background(), signedTx);
          if  err == nil {
            fmt.Println("Successfully send transaction nonce:", nonce  , ", to host: ", conn.Name )
            return nil
          }
         err_msg := err.Error()
         fmt.Println("Error in sending transaction with nonce:", nonce  , ", to host: ", conn.Name ,", error: ", err_msg)
         if strings.Contains(err_msg, "connection refused") {
             fmt.Println("Cannot connect server:", conn.Name,", deactive server. Try with others")
             connection.Pool.DeactiveNode(conn.Name)
             continue
         }
         if strings.Contains(err_msg, "insufficient funds") {
             fmt.Println("Account not enough Ether to run transaction")
             return err
         }
         retry = retry + 1
      }
      return nil
}
func (r *WalletAccount) CallContract(msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
      retry := 0
      for retry < 3 {
          conn := connection.Pool.GetConnection()
          if conn == nil {
             return nil, errors.New("Not find connection. Please check")
          }

          //Lock rpc connection before send
          conn.Mux.Lock()
          defer 	conn.Mux.Unlock()
          bs, err := conn.Client.CallContract(context.Background(),msg,blockNumber);
          if  err == nil {
            fmt.Println("Successfully send message to host: ", conn.Name )
            return bs, nil
          }
         err_msg := err.Error()
         if strings.Contains(err_msg, "connection refused") {
             fmt.Println("Cannot connect server:", conn.Name,", deactive server. Try with others")
             connection.Pool.DeactiveNode(conn.Name)
             retry = retry + 1
             continue
         }
         return bs, err
      }
      return nil, nil
}

func (r *WalletAccount) CodeAt(account common.Address, blockNumber *big.Int) ([]byte, error) {
      retry := 0
      for retry < 3 {
          conn := connection.Pool.GetConnection()
          if conn == nil {
             return  nil, errors.New("Not find connection. Please check")
          }
          //Lock rpc connection before send
          conn.Mux.Lock()
          defer 	conn.Mux.Unlock()
          bs, err := conn.Client.CodeAt(context.Background(),account,blockNumber);
          if  err == nil {
            fmt.Println("Successfully send message to host: ", conn.Name )
            return bs,err
          }
          err_msg := err.Error()
          if strings.Contains(err_msg, "connection refused") {
             fmt.Println("Cannot connect server:", conn.Name,", deactive server. Try with others")
             connection.Pool.DeactiveNode(conn.Name)
             retry = retry + 1
             continue
          }
          return bs, err
      }
      return nil, nil
}
func (r *WalletAccount) NonceAt(account common.Address) (uint64, error) {
    nonce := uint64(0)
    for _, node := range connection.Pool.Nodes {
         conn := node.GetConnection()

         conn.Mux.Lock()
         defer 	conn.Mux.Unlock()
         n, err := conn.Client.NonceAt(context.Background(),account,nil)
         if err != nil {
           continue
         }
         if n > nonce {
           nonce = n
         }
    }
    if nonce == 0 {
      return 0, errors.New("Cannot find pending nonce")
    }
    return nonce, nil
}
func (r *WalletAccount) PendingNonceAt(account common.Address) (uint64, error) {
    retry := 0
    for retry < 3 {
        conn := connection.Pool.GetConnection()
        if conn == nil {
           return  0, errors.New("Not find connection. Please check")
        }
        //Lock rpc connection before send
        conn.Mux.Lock()
        defer 	conn.Mux.Unlock()
        n, err := conn.Client.PendingNonceAt(context.Background(),account)
        if  err == nil {
          fmt.Println("Successfully get nonce from: ", conn.Name )
          return n, err
        }
        err_msg := err.Error()
        if strings.Contains(err_msg, "connection refused") {
           fmt.Println("Cannot connect server:", conn.Name,", deactive server. Try with others")
           connection.Pool.DeactiveNode(conn.Name)
           retry = retry + 1
           continue
        }
        //Other error
        return 0, err
    }
    return 0, errors.New("Cannot find nonce from system")
}
func (r *WalletAccount) MaxPendingNonceAt(account common.Address) (uint64, error) {
    nonce := uint64(0)
    for _, node := range connection.Pool.Nodes {
         conn := node.GetConnection()

         conn.Mux.Lock()
         defer 	conn.Mux.Unlock()
         n, err := conn.Client.PendingNonceAt(context.Background(),account)
         if err != nil {
           continue
         }
         if n > nonce {
           nonce = n
         }
    }
    if nonce == 0 {
      return 0, errors.New("Cannot find pending nonce")
    }
    return nonce, nil
}

func (r *WalletAccount) SubmitTransaction(signedTx *types.Transaction, nonce uint64) (error) {
    fmt.Println("Submit Transaction directly to miner ")
    start := time.Now().UnixNano()

    err :=  r.SendTransaction(signedTx, nonce)
    if err != nil {
        fmt.Println("Error send transaction", nonce," error:", err)
        return err
    }
    end := time.Now().UnixNano()
    diff:= (end-start)/1000
    fmt.Println("End Submit transaction: ", nonce,", Time: ", diff)

    return nil
}

func (w *WalletAccount) BalanceAt(account common.Address) (*big.Float, error) {
  retry := 0
  for retry < 3 {
      conn := connection.Pool.GetConnection()
      if conn == nil {
         return  nil, errors.New("Not find connection. Please check")
      }
      //Lock rpc connection before send
      conn.Mux.Lock()
      defer 	conn.Mux.Unlock()

      balance, err := conn.Client.BalanceAt(context.Background(), account, nil)
      if  err == nil {
        fmt.Println("Successfully send message to host: ", conn.Name )
        fbalance := new(big.Float)
      	fbalance.SetString(balance.String())
      	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
        return ethValue, err
      }

      err_msg := err.Error()
      if strings.Contains(err_msg, "connection refused") {
         fmt.Println("Cannot connect server:", conn.Name,", deactive server. Try with others")
         connection.Pool.DeactiveNode(conn.Name)
         retry = retry + 1
         continue
      }
      return nil, err
  }
  return nil, nil
}
