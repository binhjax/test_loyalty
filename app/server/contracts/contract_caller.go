package contracts
import (
  // "time"
  "math/big"
  // "strings"
  "fmt"
  // "encoding/json"
  "errors"
  "strings"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/common"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/jinzhu/gorm"
  "encoding/hex"
  "github.com/binhnt-teko/test_loyalty/app/server/contracts"
  "sync"
  "math"
  "strconv"
  "bytes"
)

//// ####################################### Blockchain call function ################
func  (fw *WalletHandler) CreditHistory() []string {
  retry := 0
  for retry < 3 {
    conn := fw.Client.GetConnection()
    ret, err := fw.creditHistory(conn)
    if err == nil {
       return ret
    }
    if !isConnectionError(err) {
        return ret
    }
    fw.Client.DeactiveNode(conn.Name)
    retry = retry + 1
  }
  fmt.Println("Failed to retry call creditHistory 3 times. ")
  return []string{}
}

func (fw *WalletHandler)  creditHistory(conn *RpcConnection) ([]string, error) {
  ret := []string{}
  instance, err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
  owner := common.HexToAddress("0x"+ cfg.F5Contract.Owner)

  n_credit, err := instance.GetCreditHistoryLength(&bind.CallOpts{From: owner})
  if err != nil {
    fmt.Println("Cannot Get length of wallets, error: ",err)
    return ret,err
  }
  i := int64(0)
  for i < n_credit.Int64() {
      creditId, err := instance.CreditIdx(&bind.CallOpts{From: owner},big.NewInt(i))
      i = i + 1
      if(err != nil) {
         fmt.Println("Error get Debit Idx: ", err)
         continue
      }
      creditTx, err := instance.Credits(&bind.CallOpts{From: owner},creditId)
      if(err != nil) {
         fmt.Println("Error get Debit Transactions: ", string(creditId[:])," Error: ", err)
         continue
      }
      list := []string{
        byte32ToString(creditTx.TxRef),
        byte32ToString(creditTx.StashName),
        creditTx.Amount.String(),
        creditTx.Timestamp.String(),
      }
      credit_tx :=  strings.Join(list, ",")
      ret = append(ret,credit_tx)
  }
  return ret, nil
}

func  (fw *WalletHandler) DebitHistory() []string {
  retry := 0
  for retry < 3 {
    conn := fw.Client.GetConnection()
    ret, err := fw.debitHistory(conn)
    if err == nil {
       return ret
    }
    if !isConnectionError(err) {
        return ret
    }
    fw.Client.DeactiveNode(conn.Name)
    retry = retry + 1
  }
  fmt.Println("Failed to retry call debitHistory 3 times. ")
  return []string{}
}

func  (fw *WalletHandler) debitHistory(conn *RpcConnection) ([]string,error) {
  ret := []string{}
  instance, err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)

  owner := common.HexToAddress("0x"+ cfg.F5Contract.Owner)
  n_debit, err := instance.GetDebitHistoryLength(&bind.CallOpts{From: owner})
  if err != nil {
    fmt.Println("Cannot Get length of wallets, error: ",err)
    return ret,err
  }
  i := int64(0)
  for i < n_debit.Int64() {
      debitId, err := instance.DebitIdx(&bind.CallOpts{From: owner},big.NewInt(i))
      i = i + 1
      if(err != nil) {
         fmt.Println("Error get Debit Idx: ", err)
         continue
      }
      debitTx, err := instance.Debits(&bind.CallOpts{From: owner},debitId)
      if(err != nil) {
         fmt.Println("Error get Debit Transactions: ", string(debitId[:])," Error: ", err)
         continue
      }
      list := []string{
        byte32ToString(debitTx.TxRef),
        byte32ToString(debitTx.StashName),
        debitTx.Amount.String(),
        debitTx.Timestamp.String(),
      }
      debit_tx :=  strings.Join(list, ",")
      ret = append(ret,debit_tx)
  }
  return ret,nil
}
func  (fw *WalletHandler) TransferHistory() []string {
  retry := 0
  for retry < 3 {
    conn := fw.Client.GetConnection()
    ret, err := fw.transferHistory(conn)
    if err == nil {
       return ret
    }
    if !isConnectionError(err) {
        return ret
    }
    fw.Client.DeactiveNode(conn.Name)
    retry = retry + 1
  }
  fmt.Println("Failed to retry call transferHistory 3 times. ")
  return []string{}
}
func  (fw *WalletHandler) transferHistory(conn *RpcConnection) ([]string,error) {
  ret := []string{}

  instance, err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)

  owner := common.HexToAddress("0x"+ cfg.F5Contract.Owner)
  n_transfer, err := instance.GetTransferHistoryLength(&bind.CallOpts{From: owner})
  if err != nil {
    fmt.Println("Cannot Get length of wallets, error: ",err)
    return ret, err
  }
  i := int64(0)
  for i < n_transfer.Int64() {
      transferId, err := instance.TransferIdx(&bind.CallOpts{From: owner},big.NewInt(i))
      i = i + 1
      if(err != nil) {
         fmt.Println("Error get Debit Idx: ", err)
         continue
      }
      transferTx, err := instance.Transfers(&bind.CallOpts{From: owner},transferId)
      if(err != nil) {
         fmt.Println("Error get Debit Transactions: ", string(transferId[:])," Error: ", err)
         continue
      }
      list := []string{
        byte32ToString(transferTx.TxRef),
        byte32ToString(transferTx.Sender),
        byte32ToString(transferTx.Receiver),
        transferTx.Amount.String(),
        transferTx.Note,
        strconv.Itoa(int(transferTx.TxType)),
        transferTx.Timestamp.String(),
      }
      transfer_tx :=  strings.Join(list, ",")
      ret = append(ret,transfer_tx)
  }
  return ret,nil
}

func  (fw *WalletHandler) StashNames() []string {
  retry := 0
  for retry < 3 {
    conn := fw.Client.GetConnection()
    ret, err := fw.stashNames(conn)
    if err == nil {
       return ret
    }
    if !isConnectionError(err) {
        return ret
    }
    fw.Client.DeactiveNode(conn.Name)
    retry = retry + 1
  }
  fmt.Println("Failed to retry call transferHistory 3 times. ")
  return []string{}
}

func  (fw *WalletHandler) stashNames(conn *RpcConnection) ([]string,error) {
  ret := []string{}

  instance, err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)

  owner := common.HexToAddress("0x"+ cfg.F5Contract.Owner)
  n_wallet, err := instance.GetStashNamesLenght(&bind.CallOpts{From: owner})
  if err != nil {
    fmt.Println("Cannot Get length of wallets, error: ",err)
    return ret,err
  }
  i := int64(0)
  for i < n_wallet.Int64() {
      stash_name, err := instance.StashNames(&bind.CallOpts{From: owner},big.NewInt(i))
      i = i + 1
      if(err != nil) {
         fmt.Println("Error get StashNames: ", err)
         continue
      }
      bal, err := instance.GetBalance(&bind.CallOpts{From: owner},stash_name)
      if(err != nil) {
         fmt.Println("Error get balance of: ", string(stash_name[:]), "Error:", err)
         continue
      }
      state, err := instance.GetState(&bind.CallOpts{From: owner},stash_name)
      if(err != nil) {
         fmt.Println("Error get state of: ", string(stash_name[:]), "Error:", err)
         continue
      }
      list := []string{
         byte32ToString(stash_name),
         bal.String(),
         strconv.Itoa(int(state)),
      }
      wallet_info := strings.Join(list,",")
      ret = append(ret,wallet_info)
  }
  return ret,nil
}


func  (fw *WalletHandler)  GetSummary() (int16,*big.Int, *big.Int, *big.Int,*big.Int)   {
  retry := 0
  for retry < 3 {
    conn := fw.Client.GetConnection()
    n_account, n_wallet, n_credit, n_debit, n_transfer, err := fw.getSummary(conn)
    if err == nil {
       return n_account, n_wallet, n_credit, n_debit, n_transfer
    }
    if !isConnectionError(err) {
        return n_account, n_wallet, n_credit, n_debit, n_transfer
    }
    fw.Client.DeactiveNode(conn.Name)
    retry = retry + 1
  }
  fmt.Println("Failed to retry call transferHistory 3 times. ")
  return 0,nil,nil,nil,nil
}

func  (fw *WalletHandler) getSummary(conn *RpcConnection) (int16,*big.Int, *big.Int, *big.Int,*big.Int, error)   {
      instance, err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)

      owner := common.HexToAddress("0x"+ cfg.F5Contract.Owner)
      n_account, err := instance.GetRegistedAccEthLength(&bind.CallOpts{From: owner})

      if err != nil {
        fmt.Println("Cannot Get Registed Acc Eth Length error: ",err, ", Contract: ", fw.ContractAddress.Hex())

        return 0, nil, nil, nil, nil, err
      }
      n_wallet, err := instance.GetStashNamesLenght(&bind.CallOpts{From: owner})
      if err != nil {
        fmt.Println("Cannot Get length of wallets, error: ",err)
        return n_account, nil, nil, nil, nil,err
      }
      n_credit, err := instance.GetCreditHistoryLength(&bind.CallOpts{From: owner})
      if err != nil {
        fmt.Println("Cannot Get length of credits, error: ",err)
        return n_account, n_wallet, nil, nil, nil,err
      }
      n_debit, err := instance.GetDebitHistoryLength(&bind.CallOpts{From: owner})
      if err != nil {
        fmt.Println("Cannot Get length of debit error: ",err)
        return n_account, n_wallet, n_credit, nil, nil,err
      }
      n_transfer, err := instance.GetTransferHistoryLength(&bind.CallOpts{From: owner})
      if err != nil {
        fmt.Println("Cannot Get length of transfer, error: ",err)
        return n_account, n_wallet, n_credit, n_debit, nil,err
      }
      return  n_account, n_wallet, n_credit, n_debit, n_transfer,nil
}
func (fw *WalletHandler) GetBalance(stashName string) (*big.Int, error)  {
  retry := 0
  for retry < 3 {
    conn := fw.Client.GetConnection()
    bal, err := fw.getBalance(conn,stashName)
    if err == nil {
       return bal,nil
    }
    if !isConnectionError(err) {
        return bal,err
    }
    fw.Client.DeactiveNode(conn.Name)
    retry = retry + 1
  }
  fmt.Println("Failed to retry call transferHistory 3 times. ")
  return nil,errors.New("Connection errors")
}
func (fw *WalletHandler) getBalance(conn *RpcConnection, stashName string) (*big.Int, error)  {
    fmt.Println("WalletHandler.GetBalance: Start get balance ")
    conn.Mux.Lock()
    defer  conn.Mux.Unlock()
    session,err  := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
    if err != nil {
        fmt.Println("Cannot find F5 contract")
        return nil,err
    }
    fmt.Println("WalletHandler.GetBalance: call  GetBalance")

    owner := common.HexToAddress("0x"+ cfg.F5Contract.Owner)

    return session.GetBalance(&bind.CallOpts{From: owner},stringTo32Byte(stashName))
}

func (fw *WalletHandler) GetState(stashName string) (int8, error)  {
  retry := 0
  for retry < 3 {
    conn := fw.Client.GetConnection()
    bal, err := fw.getState(conn,stashName)
    if err == nil {
       return bal,nil
    }
    if !isConnectionError(err) {
        return bal,err
    }
    fw.Client.DeactiveNode(conn.Name)
    retry = retry + 1
  }
  fmt.Println("Failed to retry call transferHistory 3 times. ")
  return 0,errors.New("Connection errors")
}
func (fw *WalletHandler) getState(conn *RpcConnection, stashName string) (int8, error)  {
  conn.Mux.Lock()
  defer  conn.Mux.Unlock()

  session, err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
  if err != nil {
      fmt.Println("Cannot find F5 contract")
      return 0,err
  }
  owner := common.HexToAddress("0x"+ cfg.F5Contract.Owner)
  return session.GetState(&bind.CallOpts{From: owner},stringTo32Byte(stashName))
}

func (fw *WalletHandler) GetTransferHistoryLength() (*big.Int, error)  {
  retry := 0
  for retry < 3 {
    conn := fw.Client.GetConnection()
    n, err := fw.getTransferHistoryLength(conn)
    if err == nil {
       return n,nil
    }
    if !isConnectionError(err) {
        return n,err
    }
    fw.Client.DeactiveNode(conn.Name)
    retry = retry + 1
  }
  fmt.Println("Failed to retry call transferHistory 3 times. ")
  return big.NewInt(0),errors.New("Connection errors")
}
func (fw *WalletHandler) getTransferHistoryLength(conn *RpcConnection) (*big.Int, error)  {
  conn.Mux.Lock()
  defer  conn.Mux.Unlock()
  session,err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
  if err != nil {
      fmt.Println("Cannot find F5 contract")
      return nil,err
  }
  owner := common.HexToAddress("0x"+ cfg.F5Contract.Owner)
  return session.GetTransferHistoryLength(&bind.CallOpts{From:owner})
}
