package contracts
import (
    // "time"
    "math/big"
    // "strings"
    "fmt"
    // "encoding/json"
    "errors"
    // "strings"
    // "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common"
    // _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    // "github.com/jinzhu/gorm"
    // "encoding/hex"
    // "sync"
    // "math"
    // "strconv"
    // "bytes"
    "github.com/binhnt-teko/test_loyalty/app/server/config"
    "github.com/binhnt-teko/test_loyalty/app/server/helper"
    "github.com/binhnt-teko/test_loyalty/app/server/connection"
    "github.com/binhnt-teko/test_loyalty/app/server/account"
   // "github.com/binhnt-teko/test_loyalty/app/server/admin"
)

func  RegisterBatchEthToContract(requestTime int64) []string {
    // cfg := config.Configuration
    ret := []string{}
    list := account.Pool.GetAccountList()
    sublist :=  []common.Address{}
    for i,item := range list {
      sublist = append(sublist,item)
      if i > 0 && i % 5 == 0 {
        if len(sublist) > 0 {
          fmt.Println("Start register sublist: ", i/5)
          tx,err := RegisterSubListAccETH(requestTime,sublist)
          if err != nil {
             ret = append(ret, err.Error())
          } 	else {
             ret = append(ret, tx.Hash().Hex())
             fmt.Println("Transaction: ",  tx.Hash().Hex())
          }
          sublist = []common.Address{}
        }
      }
    }
    if len(sublist) > 0 {
      fmt.Println("Start register last sublist")
      tx,err := RegisterSubListAccETH(requestTime,sublist)
      if err != nil {
         ret = append(ret, err.Error())
      } 	else {
         ret = append(ret, tx.Hash().Hex())
      }
    }
    return ret
}

func  RegisterSubListAccETH(requestTime int64, listAcc []common.Address) (*types.Transaction, error) {
    fmt.Println("Start RegisterAccETH, account: ", LContract.Owner)
    cfg := config.Configuration
    account := account.Pool.GetWallet(LContract.Owner)
    if account == nil {
       fmt.Println("Cannot find active account")
       return nil, errors.New("Cannot find bugdet account")
    }
    auth := account.NewTransactor()
    gasLimit,ok := cfg.Contract.GasLimit["register"]
    if !ok {
      gasLimit = cfg.Contract.GasLimitDefault
    }
    auth.GasLimit = gasLimit

    retry := 0
    for retry < 3 {
      conn := connection.Pool.GetConnection()
      // conn := connection.Pool.GetConnectionByAccount(account.Address)
      tx, err := registerAccETH(conn,auth,listAcc)
      if err == nil {
         // redisCache.LogStart(tx.Hash().Hex(), 0, requestTime )
         return tx,nil
      }
      if !helper.IsConnectionError(err) {
          return tx,err
      }
      connection.Pool.DeactiveNode(conn.Name)
      retry = retry + 1
    }
    fmt.Println("End RegisterAccETH: retry failed ")
    return nil, errors.New("Cannot find wallet in pool to create transaction")
}
func  registerAccETH(conn *connection.RpcConnection, auth *bind.TransactOpts, listAcc []common.Address) (*types.Transaction, error) {
      session,err := ContractInstance(LContract.Address,conn.Client)
      if err != nil {
          fmt.Println("Cannot find F5 contract")
          return nil,err
      }
      conn.Mux.Lock()
      defer  conn.Mux.Unlock()

      return session.RegisterAccETH(auth,listAcc)
}


func  CreateStash(requestTime int64, stashName string, typeStash int8) (*types.Transaction, error)  {
    cfg := config.Configuration
    retry := 0
    for retry <10 {
        account := account.Pool.GetAccountEth()
        if account.IsAvailable() {
          conn := connection.Pool.GetConnection()
          // conn := connection.Pool.GetConnectionByAccount(account.Address)
          session, err := ContractInstance(LContract.Address,conn.Client)
          if err != nil {
              fmt.Println("Cannot find F5 contract")
              return nil,err
          }
          auth := account.NewTransactor()
          gasLimit,ok := cfg.Contract.GasLimit["create"]
          if !ok {
            gasLimit = cfg.Contract.GasLimitDefault
          }
          auth.GasLimit = gasLimit

          conn.Mux.Lock()
          defer  conn.Mux.Unlock()
          bs := helper.StringTo32Byte(stashName)
          fmt.Println("Using: ", account, " to create Wallet: ",stashName, " len: ", len(bs) )
          tx,err := session.CreateStash(auth,bs, typeStash)
          if(err == nil){
            //Log transaction
            // redisCache.LogStart(tx.Hash().Hex(), 0, requestTime )
            return tx, err
          } else {
              fmt.Println("Error in create stash: ", err )
          }
          if helper.IsConnectionError(err) {
              connection.Pool.DeactiveNode(conn.Name)
          }
        }
        retry = retry + 1
    }
    return nil, errors.New("Cannot find wallet in pool to create transaction")
}

func  SetState(requestTime int64, stashName string, stashState int8 ) (*types.Transaction, error)  {
  cfg := config.Configuration
  retry := 0
  for retry <10 {
      account := account.Pool.GetAccountEth()
      if account.IsAvailable() {
          auth := account.NewTransactor()
          gasLimit,ok := cfg.Contract.GasLimit["state"]
          if !ok {
            gasLimit = cfg.Contract.GasLimitDefault
          }
          auth.GasLimit = gasLimit
          conn := connection.Pool.GetConnection()
          // conn := connection.Pool.GetConnectionByAccount(account.Address)
          session, err  := ContractInstance(LContract.Address,conn.Client)
          if err != nil {
              fmt.Println("Cannot find F5 contract")
              return nil,err
          }
          conn.Mux.Lock()
          defer  conn.Mux.Unlock()
          tx,err := session.SetState(auth, helper.StringTo32Byte(stashName),stashState)
          if(err == nil){
            //Log transaction
            // redisCache.LogStart(tx.Hash().Hex(), 0, requestTime )
            return tx, err
          }
          if helper.IsConnectionError(err) {
              connection.Pool.DeactiveNode(conn.Name)
          }
      }
        retry = retry + 1
  }
  return nil, errors.New("Cannot find wallet in pool to create transaction")
}

func  Debit(requestTime int64, txRef string, stashName string, amount *big.Int) (*types.Transaction, error) {
    cfg := config.Configuration
    retry := 0
    for retry <10 {
        account := account.Pool.GetAccountEth()
        if account.IsAvailable() {
            auth := account.NewTransactor()
            gasLimit,ok := cfg.Contract.GasLimit["debit"]
            if !ok {
              gasLimit = cfg.Contract.GasLimitDefault
            }
            auth.GasLimit = gasLimit
            conn := connection.Pool.GetConnection()
            // conn := connection.Pool.GetConnectionByAccount(account.Address)
            session,err := ContractInstance(LContract.Address,conn.Client)
            if err != nil {
                fmt.Println("Cannot find F5 contract")
                return nil,err
            }
            conn.Mux.Lock()
            defer  conn.Mux.Unlock()
            tx,err := session.Debit(auth, helper.StringTo32Byte(txRef),helper.StringTo32Byte(stashName),amount)
            if(err == nil){
              //Log transaction
              // redisCache.LogStart(tx.Hash().Hex(), 0, requestTime )
              return tx, err
            }
            if helper.IsConnectionError(err) {
                connection.Pool.DeactiveNode(conn.Name)
            }
        }
          retry = retry + 1
    }
    return nil, errors.New("Cannot find wallet in pool to create transaction")
}

func  Credit(requestTime int64, txRef string, stashName string, amount *big.Int) (*types.Transaction, error) {
  cfg := config.Configuration
  retry := 0
  for retry <10 {
      account := account.Pool.GetAccountEth()
      if account.IsAvailable() {
          auth := account.NewTransactor()
          gasLimit,ok := cfg.Contract.GasLimit["credit"]
          if !ok {
            gasLimit = cfg.Contract.GasLimitDefault
          }
          auth.GasLimit = gasLimit
          conn := connection.Pool.GetConnection()
          // conn := connection.Pool.GetConnectionByAccount(account.Address)
          session,err := ContractInstance(LContract.Address,conn.Client)

          if err != nil {
              fmt.Println("Cannot find F5 contract")
              return nil,err
          }
          conn.Mux.Lock()
          defer  conn.Mux.Unlock()
          tx,err :=  session.Credit(auth, helper.StringTo32Byte(txRef),helper.StringTo32Byte(stashName),amount)
          if(err == nil){
            //Log transaction
            // redisCache.LogStart(tx.Hash().Hex(), 0, requestTime )
            return tx, err
          }
          if helper.IsConnectionError(err) {
              connection.Pool.DeactiveNode(conn.Name)
          }
      }
  }
  return nil, errors.New("Cannot find wallet in pool to create transaction")
}

func  Transfer(requestTime int64, txRef string, sender string, receiver string, amount *big.Int, note string, txType int8) (*types.Transaction, error) {
  cfg := config.Configuration
  retry := 0
  for retry <10 {
      account := account.Pool.GetAccountEth()
      if account.IsAvailable() {
          auth := account.NewTransactor()
          gasLimit,ok := cfg.Contract.GasLimit["transfer"]
          if !ok {
            gasLimit = cfg.Contract.GasLimitDefault
          }
          auth.GasLimit = gasLimit
          conn := connection.Pool.GetConnection()
          // conn := connection.Pool.GetConnectionByAccount(account.Address)
          session,err := ContractInstance(LContract.Address,conn.Client)
          if err != nil {
              fmt.Println("Cannot find F5 contract")
              return nil,err
          }
          conn.Mux.Lock()
          defer  conn.Mux.Unlock()
          tx, err :=  session.Transfer(auth, helper.StringTo32Byte(txRef),helper.StringTo32Byte(sender),helper.StringTo32Byte(receiver),amount,note,txType)
          if(err == nil){
            //Log transaction
            // redisCache.LogStart(tx.Hash().Hex(), 0, requestTime )
            return tx, err
          }
          if helper.IsConnectionError(err) {
              connection.Pool.DeactiveNode(conn.Name)
          }

      }
        retry = retry + 1
  }
  return nil, errors.New("Cannot find wallet in pool to create transaction")
}
