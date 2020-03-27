package controllers

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	// "github.com/valyala/fasthttp"
  // "github.com/binhnt-teko/test_loyalty/app/server/middleware"
  "github.com/binhnt-teko/test_loyalty/app/server/account"
	// "github.com/binhnt-teko/test_loyalty/app/server/config"
	"time"
	// "encoding/json"
)

type ApiWallet struct {}

//Load account from keystore
func (api ApiWallet) Load(c  *routing.Context) error {
  	if err := account.Pool.LoadAccounts(); err != nil {
  			fmt.Println(time.Now()," Cannot load wallet from redis ", err)
        return Response(c,99,"Cannot load wallet from redis ",err)
  	}
    return Response(c,0,"Load wallet successfully",nil)
}
//Load account from file
func (api ApiWallet) List(c  *routing.Context) error {
    wallets := account.Pool.GetAccountList()
		return Response(c,0,"Wallet List",wallets)
}
//List account in redis
func (api ApiWallet) New(c  *routing.Context) error {
	  account, err  := account.Pool.NewWalletAccount()
		if err != nil {
      fmt.Println(" Cannot create new account ",err)
      return Response(c, 99 ," Cannot create new account",err)
    }
    return Response(c,0,"New Account",account)
}
func (api ApiWallet) Save(c  *routing.Context) error {
	err  := account.Pool.Save()
	if err != nil {
		fmt.Println(" Cannot save wallet to redis ",err)
		return Response(c, 99 ," Cannot save wallet to redist",err)
	}
	return Response(c,0,"Successfully save wallets to redis",nil)
}
func (api ApiWallet) FillGas(c  *routing.Context) error {
	txs  := account.Pool.AutoFillGas()
	return Response(c,0,"Transactions:",txs)
}
func (api ApiWallet) Balance(c  *routing.Context) error {
	address := c.Param("address")
	wallet := account.Pool.GetWallet(address)
	if wallet != nil  {
		   bal, err := wallet.EthBalaneOf()
			 if err != nil {
					fmt.Println("Cannot get wallet balance")
					return Response(c,99,"Cannot get wallet balance:",err)
			 }
			 ba,_ := bal.Float64()
		 	 return Response(c,0,"Balance of " + address,ba)
	}
	return Response(c,99,"Cannot find wallet",nil)
}
