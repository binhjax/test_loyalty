package controllers

import (
		"fmt"
		"github.com/qiangxue/fasthttp-routing"
		// "github.com/valyala/fasthttp"
	  "github.com/binhnt-teko/test_loyalty/app/server/admin"
	  // "github.com/binhnt-teko/test_loyalty/app/server/config"
		// "os"
	  // "io/ioutil"
		// "gopkg.in/yaml.v2"
			"encoding/json"
			"time"
)

type ApiAdmin struct {}

// swagger:route POST /api/admin/network/addpeers admin_addpeers addPeersEndpoint
// Addpeers in all node of ethereum blockchain .
// responses:
//   200: commonResponse

func (api ApiAdmin) AddPeers(c *routing.Context) error {
	  if err := admin.AddPeers(); err != nil {
				return Response(c,98,"Add peer error ",err)
		}
		return Response(c,0,"Add peers successfully", nil)
}

// swagger:route POST /api/admin/network/pool admin_pool poolEndpoint
// Get current pool of one geth  .
// responses:
//   200: commonResponse

func (api ApiAdmin) GetPool(c *routing.Context) error {
		return Response(c,0,"Data", nil)
}


// swagger:route POST /api/admin/transaction admin_transaction transactionEndpoint
// Get current pool of one geth  .
// responses:
//   200: commonResponse

func (api ApiAdmin) GetTransactionByHash(c *routing.Context) error {
	hash := c.Param("hash")
  receipt, err := admin.GetTransactionState(hash)
	if err != nil {
		return Response(c,98,"get transation error ",err)
	}
	return Response(c,0,"Data", receipt)
}

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
  log, err := admin.GetLog(filterLog.FromBlock, filterLog.ToBlock, filterLog.Addresses )
	if err != nil {
		return Response(c,98,"get transation error ",err)
	}
	return Response(c,0,"Data", log)
}
