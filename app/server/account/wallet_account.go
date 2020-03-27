package account

import (
    "github.com/binhnt-teko/test_loyalty/app/server/config"
    "math/big"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "errors"
    "strings"
    "fmt"
    "crypto/ecdsa"
    "encoding/hex"
    "math"
    "sync"
    "sync/atomic"
)

type WalletAccount struct {
    Address string
    Nonce uint64
    Active bool
    PrivateKey *ecdsa.PrivateKey
    Account *config.TokenAccount
    Mutex sync.Mutex
}

func  (w *WalletAccount) GetPrivateKey() string {
     return hex.EncodeToString(crypto.FromECDSA(w.PrivateKey))
}

func (w *WalletAccount) IsAvailable() bool {
   return w.Active
}
func  (w *WalletAccount) NewTransactor() *bind.TransactOpts {
      cfg := config.Configuration
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
     gasPrice.SetString(cfg.Contract.GasPrice,10)
     auth.GasPrice = gasPrice
     auth.GasLimit = cfg.Contract.GasLimitDefault
     return auth
}

func (w *WalletAccount) EthTransfer(to string, amount string) (string,uint64,error)  {
      cfg := config.Configuration
      //1. Get nonce
      nonce := w.GetNonce()

      //2. Prepare transaction
      gLimit := cfg.Contract.GasLimitDefault
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
        fmt.Println(" Cannot sign contxract: ", err)
        return "",0,err
      }

      signedTx, err := rawTx.WithSignature(signer, signature)

      txhash := strings.TrimPrefix(signedTx.Hash().Hex(),"0x")

      err = w.SubmitTransaction(signedTx,nonce)

      return txhash, nonce, err
}


func (w *WalletAccount) EthBalaneOf() (*big.Float, error) {
    w.Mutex.Lock()
    defer w.Mutex.Unlock()
    account := common.HexToAddress("0x" + w.Address)
    return w.BalanceAt(account)
}

func (w *WalletAccount) UpdateNonce(nonce uint64)  {
    fmt.Println("Update nonce of: ",w.Address," Nonce:",nonce)
    atomic.StoreUint64(&w.Nonce, nonce-1)
}

func (w *WalletAccount) GetNonce() uint64 {
    keyAddr := common.HexToAddress(w.Address)
    nonce, err := w.PendingNonceAt(keyAddr)
    if err != nil {
      fmt.Errorf("failed to retrieve account nonce: %v", err)
      w.UpdateNonce(0)
      return 0
    }
    fmt.Println("Sync nonce from eth: ",nonce)
    w.UpdateNonce(nonce)
    // nonce = atomic.AddUint64(&w.Nonce, 1)
    fmt.Println("WalletAccount.GetNonce: Get Nonce:",nonce)
    return nonce
}
