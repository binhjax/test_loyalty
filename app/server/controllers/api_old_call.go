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
	"time"
	"fmt"
	"encoding/json"
  "github.com/binhnt-teko/test_loyalty/app/server/contracts"
)


// swagger:route POST /api/admin/log admin_viewlog logEndpoint
// Get blockchain log
// responses:
//   200: commonResponse

func (api ApiAdmin) GetLog(c *routing.Context) error {
	// fmt.Println("Start getlog")
	var filterLog = FilterLog{}
	val := c.PostBody()
	err := json.Unmarshal([]byte(val), &filterLog)
	if err != nil {
			fmt.Println(time.Now()," Cannot parse filterlog ", err)
			return err
	}
  log, err := contracts.GetLog(filterLog.FromBlock, filterLog.ToBlock, filterLog.Addresses )
	if err != nil {
		return Response(c,98,"get transation error ",err)
	}
	return Response(c,0,"Data", log)
}


//1. CreditHistory
func (api ApiTest) CreditHistory(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  txs := contracts.CreditHistory();
  return Response(c,0,"CreditHistory ", txs)
}
//1. CreditHistory
func (api ApiTest) DebitHistory(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  txs := contracts.DebitHistory();
  return Response(c,0,"DebitHistory ", txs)
}
func (api ApiTest) TransferHistory(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  txs := contracts.TransferHistory();
  return Response(c,0,"TransferHistory ", txs)
}
func (api ApiTest) StashNames(c *routing.Context) error {
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  txs := contracts.StashNames();
  return Response(c,0,"StashNames ", txs)
}
func (api ApiTest) StashBalance(c *routing.Context) error {
	stashName := c.Param("name")
  if !contracts.Loaded {
      return Response(c,99,"Contract is not loaded. Pls load ", nil)
  }
	// requestTime := time.Now().UnixNano()
  amount, err := contracts.GetBalance(stashName);
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
  state, err := contracts.GetState(stashName);
	if err != nil {
		return Response(c,0,"Cannot get balance ", err)
	}
  return Response(c,0,"state ", state)
}
