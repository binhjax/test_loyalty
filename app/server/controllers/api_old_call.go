package controllers

import (
	// "fmt"
	"github.com/qiangxue/fasthttp-routing"
	// "github.com/valyala/fasthttp"
  // "github.com/binhnt-teko/test_loyalty/app/server/middleware"
  // "github.com/binhnt-teko/test_loyalty/app/server/helper"
	// "github.com/binhnt-teko/test_loyalty/app/server/workers"
  // "github.com/binhnt-teko/test_loyalty/app/server/config"
	 // "math/big"
	// "time"
	// "fmt"
	// "encoding/json"
  "github.com/binhnt-teko/test_loyalty/app/server/contracts"
)

//1. CreditHistory
func (api ApiTest) CreditHistory(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  txs := contracts.LContract.CreditHistory();
  return Response(c,0,"CreditHistory ", txs)
}
//1. CreditHistory
func (api ApiTest) DebitHistory(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  txs := contracts.LContract.DebitHistory();
  return Response(c,0,"DebitHistory ", txs)
}
func (api ApiTest) TransferHistory(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  txs := contracts.LContract.TransferHistory();
  return Response(c,0,"TransferHistory ", txs)
}
func (api ApiTest) StashNames(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  txs := contracts.LContract.StashNames();
  return Response(c,0,"StashNames ", txs)
}
func (api ApiTest) StashBalance(c *routing.Context) error {
	stashName := c.Param("name")
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  amount, err := contracts.LContract.GetBalance(stashName);
	if err != nil {
		return Response(c,0,"Cannot get balance ", err)
	}
  return Response(c,0,"balance ", amount)
}
func (api ApiTest) StashGetState(c *routing.Context) error {
	stashName := c.Param("name")
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  state, err := contracts.LContract.GetState(stashName);
	if err != nil {
		return Response(c,0,"Cannot get balance ", err)
	}
  return Response(c,0,"state ", state)
}
