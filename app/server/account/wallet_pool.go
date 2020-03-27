package account

import (
    "github.com/binhnt-teko/test_loyalty/app/server/config"
    "math/big"
    "fmt"
    "errors"
    "strings"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/ethereum/go-ethereum/core/types"
    "encoding/hex"
    "sync"
    "math"
    "strconv"
    "time"
    "encoding/json"
)

type WalletPool struct {
    Wallets []*WalletAccount
    Current int
    Mutex sync.Mutex
}

var Pool *WalletPool

//// ####################################### Processing Support function ################
func Init(){
      Pool =  &WalletPool{
        Current: 0,
      }
      Pool.LoadAccounts()
      // Pool.AutoFillGas()
}

func (wp *WalletPool) GetAccountEth() *WalletAccount{
    wp.Mutex.Lock()
    defer wp.Mutex.Unlock()
    len := len(wp.Wallets)
    if len == 0 {
      return nil
    }
    if wp.Current >= len {
         wp.Current = wp.Current % len
    }
    wallet := wp.Wallets[wp.Current]
    wp.Current = wp.Current + 1
    return wallet
}

func (wp *WalletPool) LoadAccounts() error {
  tokens, err  := config.Redis.Keys("token.*").Result()
  if err != nil {
    // handle error
    fmt.Println(" Cannot get addresses ")
    return err
  }

  accounts := []config.TokenAccount{}
  for _, token := range tokens {
    val, err := config.Redis.Get(token).Result()
    if err != nil {
        fmt.Println(time.Now()," Cannot find token: ", token)
        continue
    }
    account := config.TokenAccount{}
    err = json.Unmarshal([]byte(val), &account)
    if err != nil {
        fmt.Println(time.Now()," Cannot parse account ", err)
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
  return nil
}

func (wp *WalletPool) GetAccountList() ([]common.Address) {
   fmt.Println("WalletAccount.GetAccountList: start read wallets")
   wp.Mutex.Lock()
   defer wp.Mutex.Unlock()
   accounts := []common.Address{}
   for _,wallet := range wp.Wallets {
       if wallet.Active {
         address := common.HexToAddress("0x"+wallet.Address)
         accounts = append(accounts,address)
       }
   }
   fmt.Println("WalletAccount.GetAccountList: end read wallets:",len(accounts))
   return accounts
}

func (wp *WalletPool) GetWallet(addr string) *WalletAccount {
    for _, wallet := range wp.Wallets {
       if wallet.Address == addr {
         return wallet
       }
    }
    return nil
}

func (wp *WalletPool) EthTransfer(from string,to string,amount string) (string,error) {
   cfg := config.Configuration
   wallet := wp.GetWallet(from)

   fromAddress := common.HexToAddress("0x" + wallet.Address)
   nonce, err := wallet.PendingNonceAt(fromAddress)
   if err != nil {
     fmt.Println("Error in getting nonce ")
     return "", err
   }

   gLimit := cfg.Contract.GasLimitDefault
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
   err = wallet.SubmitTransaction(signedTx,nonce)

   return txhash, err
}

func (wp *WalletPool) AutoFillGas() []string {
    wp.Mutex.Lock()
    defer wp.Mutex.Unlock()

    cfg := config.Configuration
    ret := []string{}
    for _, wallet := range wp.Wallets {
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

           txhash, err := wp.EthTransfer(cfg.Contract.EthBudget, wallet.Address,strconv.Itoa(fill_account))
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

func (wp *WalletPool) NewWalletAccount() (string, error) {
    wp.Mutex.Lock()
    defer wp.Mutex.Unlock()
    privateKey, err := crypto.GenerateKey()
    if err != nil {
      return "",err
    }
    address := crypto.PubkeyToAddress(privateKey.PublicKey)

    account := address.Hex()
    account = strings.TrimPrefix(account,"0x")
    account = strings.ToLower(account)

    priKey :=  hex.EncodeToString(crypto.FromECDSA(privateKey))

    new_account := &config.TokenAccount{
      Address: account,
      PrivateKey: priKey,
    }
    wallet := WalletAccount{
      Address: account,
      PrivateKey: privateKey,
      Nonce: 0,
      Account: new_account,
      Active: true,
    }
    wp.Wallets = append(wp.Wallets,&wallet)
    return  address.Hex(), nil
}

func (wp *WalletPool) Save() error {
  fmt.Println("WalletAccount.GetAccountList: start read wallets")
  wp.Mutex.Lock()
  defer wp.Mutex.Unlock()
  for _,wallet := range wp.Wallets {
    account := wallet.Account
    value, err := json.Marshal(wallet.Account)
    if err != nil {
        fmt.Println("Cannot convert to json: ", err)
        continue
    }
    err = config.Redis.Set("token." + account.Address,string(value), 0).Err()
    if err != nil {
      fmt.Println("Cannot save to redis : ", err)
    }
  }
  fmt.Println("Finished save wallets to redis")
  return nil
}
