package controllers

import (
	// "fmt"
	"github.com/qiangxue/fasthttp-routing"
	// "github.com/valyala/fasthttp"
  // "github.com/binhnt-teko/test_loyalty/app/server/middleware"
  // "github.com/binhnt-teko/test_loyalty/app/server/helper"
	// "github.com/binhnt-teko/test_loyalty/app/server/workers"
  // "github.com/binhnt-teko/test_loyalty/app/server/config"
	 "math/big"
	"time"
	"fmt"
	"encoding/json"
  "github.com/binhnt-teko/test_loyalty/app/server/contracts"
)

type ApiTest struct {}

//1. Register eth
func (api ApiTest) RegisterEth(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	requestTime := time.Now().UnixNano()
  txs := contracts.RegisterBatchEthToContract(requestTime);
  return Response(c,0,"Load transaction ", txs)
}
//2. Create stash
type StashForm struct {
  Name string `json:"name"`
	Type int8 `json:"type"`
}

func (api ApiTest) CreateStash(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }

	var stash = &StashForm{}
	err := json.Unmarshal([]byte(c.PostBody()), &stash)
	if err != nil {
			fmt.Println(time.Now()," Cannot parse stash ", err)
			return Response(c,99,"Cannot parse stash",err)
	}

	requestTime := time.Now().UnixNano()
  tx, err := contracts.CreateStash(requestTime, stash.Name, stash.Type)
	if err != nil {
			fmt.Println(time.Now()," Cannot create transaction  ", err)
			return Response(c,99,"Cannot create transaction ",err)
	}
  return Response(c,0,"Load transaction ", tx.Hash().Hex())
}

func (api ApiTest) StashState(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }

	var stash = &StashForm{}
	err := json.Unmarshal([]byte(c.PostBody()), &stash)
	if err != nil {
			fmt.Println(time.Now()," Cannot parse stash ", err)
			return Response(c,99,"Cannot parse stash",err)
	}

	requestTime := time.Now().UnixNano()
  tx, err := contracts.SetState(requestTime, stash.Name, stash.Type)
	if err != nil {
			fmt.Println(time.Now()," Cannot create transaction  ", err)
			return Response(c,99,"Cannot create transaction ",err)
	}
  return Response(c,0,"Load transaction ", tx.Hash().Hex())
}

type DebitCreditForm struct {
  TxRef string `json:"txRef"`
	Name string `json:"name"`
	Amount int64 `json:"amount"`
}

func (api ApiTest) StashDebit(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }

	var stash = &DebitCreditForm{}
	err := json.Unmarshal([]byte(c.PostBody()), &stash)
	if err != nil {
			fmt.Println(time.Now()," Cannot parse request ", err)
			return Response(c,99,"Cannot parse request",err)
	}

	requestTime := time.Now().UnixNano()
  tx, err := contracts.Debit(requestTime, stash.TxRef, stash.Name, big.NewInt(stash.Amount))
	if err != nil {
			fmt.Println(time.Now()," Cannot create transaction  ", err)
			return Response(c,99,"Cannot create transaction ",err)
	}
  return Response(c,0,"Load transaction ", tx.Hash().Hex())
}

func (api ApiTest) StashCredit(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }

	var stash = &DebitCreditForm{}
	err := json.Unmarshal([]byte(c.PostBody()), &stash)
	if err != nil {
			fmt.Println(time.Now()," Cannot parse stash ", err)
			return Response(c,99,"Cannot parse stash",err)
	}

	requestTime := time.Now().UnixNano()
  tx, err := contracts.Credit(requestTime, stash.TxRef, stash.Name, big.NewInt(stash.Amount))
	if err != nil {
			fmt.Println(time.Now()," Cannot create transaction  ", err)
			return Response(c,99,"Cannot create transaction ",err)
	}
  return Response(c,0,"Load transaction ", tx.Hash().Hex())
}


type TransferForm struct {
  TxRef string `json:"txRef"`
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount int64 `json:"amount"`
	Note string `json:"note"`
	TxType int8 `json:"txType"`
}


func (api ApiTest) StashTransfer(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }

	var stash = &TransferForm{}
	err := json.Unmarshal([]byte(c.PostBody()), &stash)
	if err != nil {
			fmt.Println(time.Now()," Cannot parse stash ", err)
			return Response(c,99,"Cannot parse stash",err)
	}

	requestTime := time.Now().UnixNano()
  tx, err := contracts.Transfer(requestTime, stash.TxRef, stash.Sender,stash.Receiver, big.NewInt(stash.Amount), stash.Note, stash.TxType)
	if err != nil {
			fmt.Println(time.Now()," Cannot create transaction  ", err)
			return Response(c,99,"Cannot create transaction ",err)
	}
  return Response(c,0,"Load transaction ", tx.Hash().Hex())
}
