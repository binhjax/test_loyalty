package account

import (
    "github.com/binhnt-teko/test_loyalty/src/config"
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

type WalletPool struct {
    Wallets []*WalletAccount
    Current int
    Mutex sync.Mutex
}

var Pool *WalletPool

//// ####################################### Processing Support function ################
func Init(){
      Pool :=  &WalletPool{
        Current: 0,
      }
      WalletPool.LoadAccounts()
      WalletPool.AutoFillGas()
}

func (fw *WalletAccount) GetAccountEth() *WalletAccount{
    fw.Mutex.Lock()
    defer fw.Mutex.Unlock()
    len := len(fw.Wallets)
    if len == 0 {
      return nil
    }
    if fw.Current >= len {
         fw.Current = fw.Current % len
    }
    wallet := fw.Wallets[fw.Current]
    fw.Current = fw.Current + 1
    return wallet
}

func (wp *WalletPool) LoadAccounts(){
  tokens, err  := config.Redis.Keys("token.*").Result()
  if err != nil {
    // handle error
    fmt.Println(" Cannot get addresses ")
    return nil, err
  }

  accounts := []TokenAccount{}
  for _, token := range tokens {
    val, err := config.Redis.Get(token).Result()
    if err != nil {
        fmt.Println(time.Now()," Cannot find token: ", token)
        continue
    }
    account := &config.TokenAccount{}
    err := json.Unmarshal([]byte(val), data)
    if err != nil {
        fmt.Println(time.Now()," Cannot parse data ", err)
        continue
    }
    accounts = append(accounts,account)
  }

  wallets := []*WalletAccount{}
  for _, account := range accounts {
      fmt.Println("Load wallet: ",account.Address)
      b, err := hex.DecodeString(account.PrivateKey)
     if err != nil {
          fmt.Println("invalid hex string: " + account.PrivateKey)
         continue
     }
      privateKey := crypto.ToECDSAUnsafe(b)
      wallet := WalletAccount{
        Address: account.Address,
        PrivateKey: privateKey,
        Active: true,
        Account: &account,
        Nonce: 0,
      }
      wallets = append(wallets,&wallet)
  }
  fmt.Println("End load accounts from db: ", len(wallets))
  wp.Mutex.Lock()
  defer wp.Mutex.Unlock()
  wp.Wallets = wallets
}

func (wp *WalletPool) GetAccountList() ([]common.Address) {
   fmt.Println("WalletAccount.GetAccountList: start read wallets")
   wp.Mutex.Lock()
   defer wp.Mutex.Unlock()
   accounts := []common.Address{}
   for _,wallet := range fw.Wallets {
       if wallet.Active {
         address := common.HexToAddress("0x"+wallet.Address)
         accounts = append(accounts,address)
       }
   }
   fmt.Println("WalletAccount.GetAccountList: end read wallets:",len(accounts))
   return accounts
}

func (wp *WalletPool) GetAccountEthAddress(addr string) *WalletAccount {
    for _, wallet := range wp.Wallets {
       if wallet.Address == addr {
         return wallet
       }
    }
    return nil
}

func (wp *WalletPool) EthTransfer(from string,to string,amount string) (string,error) {
   wallet := fw.GetAccountEthAddress(from)

   fromAddress := common.HexToAddress("0x" + wallet.Address)
   nonce, err := wallet.Routing.PendingNonceAt(fromAddress)
   if err != nil {
     fmt.Println("Error in getting nonce ")
     return "", err
   }

   gLimit := cfg.Contract.GasLimit
   gPrice := cfg.Contract.GasPrice

   gasLimit := uint64(gLimit)
   gasPrice := new(big.Int)
   gasPrice, _ = gasPrice.SetString(gPrice, 10)

   toAddress := common.HexToAddress("0x" + to)

   eth_unit := big.NewFloat(math.Pow10(18))
   amount_value := new(big.Float)
   value, ok := amount_value.SetString(amount)

   if !ok {
        fmt.Println("SetString: error")
        return "", errors.New("convert amount error")
   }
   value = value.Mul(value,eth_unit)

   value_transfer := new(big.Int)
   value.Int(value_transfer)

   var data []byte
   rawTx := types.NewTransaction(nonce, toAddress, value_transfer, gasLimit, gasPrice, data)

   signer := types.FrontierSigner{}
   signature, err := crypto.Sign(signer.Hash(rawTx).Bytes(), wallet.PrivateKey)
   if err != nil {
     fmt.Println(" Cannot sign contract: ", err)
     return "",err
   }

   signedTx, err := rawTx.WithSignature(signer, signature)

   txhash := strings.TrimPrefix(signedTx.Hash().Hex(),"0x")
   err = wallet.Routing.SubmitTransaction(signedTx,nonce)

   return txhash, err
}

func (wp *WalletPool) AutoFillGas() []string {
    fw.Mutex.Lock()
    defer fw.Mutex.Unlock()

    ret := []string{}
    for _, wallet := range fw.Wallets {
      bal, err := wallet.EthBalaneOf()
      if err != nil {
         fmt.Println("Cannot get wallet balance. Deactive wallet")
         wallet.Active = false
         continue
      }
      ba,_ := bal.Float64()
      if ba < 1000 {
         fmt.Println("Create transaction to fillGass from budget")
         var fill_account int = int(1000 - ba)

         txhash, err := fw.EthTransfer(cfg.F5Contract.EthBudget, wallet.Address,strconv.Itoa(fill_account))
         if err != nil {
           fmt.Println("Cannot fill more gas. Deactive wallet ")
           wallet.Active = false
           continue
         }
         fmt.Println("Fill Eth to account: ", wallet.Address, " transaction: ", txhash)
         ret = append(ret,txhash)
      } else {
         fmt.Println("Account: ", wallet.Address, " balance: ", ba)
      }
    }
    return ret
}

//#########################   Non blockchain function ##########################
func (wp *WalletPool) NewAccountEth() (string, error) {
      privateKey, err := crypto.GenerateKey()
      if err != nil {
        return "",err
      }
      address := crypto.PubkeyToAddress(privateKey.PublicKey)

      account := address.Hex()
      account = strings.TrimPrefix(account,"0x")
      account = strings.ToLower(account)

      priKey :=  hex.EncodeToString(crypto.FromECDXSA(privateKey))

      new_account := &TokenAccount{
        Address: account,
        PrivateKey: priKey,
        Active: true,
      }

      fmt.Println("Update account to redis ")

      fmt.Println("Update account to wallet ")
      wallet := WalletAccount{
        Address: account,
        PrivateKey: privateKey,
        Nonce: 0,
        Account: new_account,
        Active: false,
      }
      fw.Wallets = append(fw.Wallets,&wallet)
      return account, nil
}
