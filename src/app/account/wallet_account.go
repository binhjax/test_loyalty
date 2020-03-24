package account

import (
  // "time"
  // "context"
  "github.com/binhnt-teko/test_loyalty/src/app/connection"
  "github.com/binhnt-teko/test_loyalty/src/config"
  "math/big"
  // "github.com/ethereum/go-ethereum"
  "github.com/ethereum/go-ethereum/core/types"
  // "github.com/ethereum/go-ethereum/accounts/abi"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "errors"
  "strings"
  "fmt"
  // "encoding/json"
  "crypto/ecdsa"
  "encoding/hex"
  // "time"
  "math"
  "github.com/jinzhu/gorm"
  "sync"
)

type TokenAccount struct {
    gorm.Model
    Address string
    WalletId string
    PrivateKey string `sql:"type:text"`
    Active bool `gorm:"default:false"`
}

type WalletAccount struct {
    Routing *connection.RpcRouting
    Address string
    Nonce uint64
    Active bool
    PrivateKey *ecdsa.PrivateKey
    Account *TokenAccount
    Mutex sync.Mutex
}

func  (w *WalletAccount) GetPrivateKey() string {
     return hex.EncodeToString(crypto.FromECDSA(w.PrivateKey))
}

func  (w *WalletAccount) NewTransactor(cfg *config.Config) *bind.TransactOpts {
      key := w.PrivateKey
    	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
    	auth := &bind.TransactOpts{
    		From: keyAddr,
    		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
    			if address != keyAddr {
    				return nil, errors.New("not authorized to sign this account")
    			}
    			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
    			if err != nil {
    				return nil, err
    			}
    			return tx.WithSignature(signer, signature)
    		},
    	}
     gasPrice := new(big.Int)
     gasPrice.SetString(cfg.F5Contract.GasPrice,10)
     auth.GasPrice = gasPrice
     auth.GasLimit = cfg.F5Contract.GasLimitDefault
     return auth
}

func (w *WalletAccount) EthTransfer(cfg *config.Config, to string, amount string) (string,uint64,error)  {
      //1. Get nonce
      nonce := w.GetNonce()

      //2. Prepare transaction
      gLimit := cfg.Contract.GasLimit
      gPrice := cfg.Contract.GasPrice

      fromAddress := common.HexToAddress("0x" + w.Address)
      fmt.Println("Address: ", fromAddress.Hex(), ",Hash: ",fromAddress.Hash())

      gasLimit := uint64(gLimit)
      gasPrice := new(big.Int)
      gasPrice, _ = gasPrice.SetString(gPrice, 10)

      toAddress := common.HexToAddress("0x" + to)

      eth_unit := big.NewFloat(math.Pow10(18))
      amount_value := new(big.Float)
      value, ok := amount_value.SetString(amount)

      if !ok {
           fmt.Println("SetString: error")
           return "", 0, errors.New("convert amount error")
      }
      value = value.Mul(value,eth_unit)

      value_transfer := new(big.Int)
      value.Int(value_transfer)

      var data []byte
      rawTx := types.NewTransaction(nonce, toAddress, value_transfer, gasLimit, gasPrice, data)

      //signer := types.NewEIP155Signer(chainID)
      //signer := types.HomesteadSigner{}
      signer := types.FrontierSigner{}
      signature, err := crypto.Sign(signer.Hash(rawTx).Bytes(), w.PrivateKey)
      if err != nil {
        fmt.Println(" Cannot sign contract: ", err)
        return "",0,err
      }

      signedTx, err := rawTx.WithSignature(signer, signature)

      txhash := strings.TrimPrefix(signedTx.Hash().Hex(),"0x")

      err = w.Routing.SubmitTransaction(signedTx,nonce)

      return txhash, nonce, err
}


func (w *WalletAccount) EthBalaneOf() (*big.Float, error) {
    w.Mutex.Lock()
    defer w.Mutex.Unlock()
    account := common.HexToAddress("0x" + w.Address)
    return w.Routing.BalanceAt(account)
}
