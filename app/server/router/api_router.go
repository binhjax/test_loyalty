package router

import (
   "github.com/qiangxue/fasthttp-routing"
   "github.com/binhnt-teko/test_loyalty/app/server/controllers"
	 "github.com/binhnt-teko/test_loyalty/app/server/middleware"
)

func NewRouter()  *routing.Router {
	router := routing.New()

  index := new(controllers.IndexController)
	apiauth := new(controllers.ApiAuth)
	// apiv1 := new(controllers.ApiV1)
  apiAdmin := new(controllers.ApiAdmin)
  apiAccount := new(controllers.ApiAccount)
  apiWallet := new(controllers.ApiWallet)
  apiContract := new(controllers.ApiContract)
  apiQueue := new(controllers.ApiQueue)
  apiTest := new(controllers.ApiTest)

  router.Get("/", index.Index)


  auth := router.Group("/auth")
  {
      auth.Post("/loadjwt", apiauth.LoadJwtAccount)
      auth.Post("/login", apiauth.Login)
  }
	authorized := router.Group("/api")

  // Authentication required
	authorized.Use(middleware.JWTMiddleware())
	{
			admin := authorized.Group("/admin")
			{
          admin.Post("/log",apiAdmin.GetLog)
          network := admin.Group("/network")
          {
             network.Get("/addpeers",apiAdmin.AddPeers)
          }
          transaction := admin.Group("/transaction")
          {
             transaction.Get("/<hash>",apiAdmin.GetTransactionByHash)
          }
          account := admin.Group("/account")
          {
             account.Post("/keystore",apiAccount.LoadKeystore)
             account.Post("/loadfile",apiAccount.LoadFile)
             account.Post("/savefile",apiAccount.SaveFile)
             account.Get("/tokens",apiAccount.ListToken)
          }

          wallet := admin.Group("/wallet")
          {
             wallet.Post("/load",apiWallet.Load)
             wallet.Post("/list",apiWallet.List)
             wallet.Post("/new",apiWallet.New)
             wallet.Get("/save",apiWallet.Save)
             wallet.Get("/fillgas",apiWallet.FillGas)
             wallet.Get("/balance/<address>",apiWallet.Balance)
          }
          contract := admin.Group("/contract")
          {
             contract.Post("/load",apiContract.Load)
             contract.Get("/view",apiContract.View)
             contract.Post("/deploy",apiContract.Deploy)
             contract.Post("/save",apiContract.Save)
             contract.Post("/code",apiContract.Code)
          }
			}
    	queue := authorized.Group("/queue")
      {
         queue.Post("/task",apiQueue.Task)
         queue.Post("/status/<taskid>",apiQueue.Status)
      }
      test := authorized.Group("/test")
      {
        test.Post("/register",apiTest.RegisterEth)
        stash := test.Group("/stash")
        {
          stash.Post("/create", apiTest.CreateStash)
          stash.Post("/state", apiTest.StashState)
          stash.Post("/debit", apiTest.StashDebit)
          stash.Post("/credit", apiTest.StashCredit)
          stash.Post("/transfer", apiTest.StashTransfer)
          stash.Get("", apiTest.StashNames)
          stash.Get("/<name>/balance", apiTest.StashBalance)
          stash.Get("/<name>/state", apiTest.StashGetState)
        }
        history := test.Group("/history")
        {
          history.Get("/credit", apiTest.CreditHistory)
          history.Get("/debit", apiTest.DebitHistory)
          history.Get("/transfer", apiTest.TransferHistory)
        }

        // cash := test.Group("/cash")
        // {
        //   cash.Put("credit/<address>/<amount>/<traceid>", apiv1.Cash_credit)
        //   cash.Put("debit/<address>/<amount>/<traceid>", apiv1.Cash_debit)
        //   cash.Put("transfer/<from>/<to>/<amount>/<note>/<traceid>", apiv1.Cash_transfer)
        // }
        // balance := test.Group("/balance")
        // {
        //   balance.Get("<address>", apiv1.Balance)
        //   balance.Get("all", apiv1.Balance_all)
        // }
        // account := test.Group("/account")
        // {
        //   account.Post("new", apiv1.Account_new)
        //   account.Get("total", apiv1.Account_total)
        //   account.Get("list/active", apiv1.Account_list_active)
        //   account.Get("list/inactive", apiv1.Account_list_inactive)
        //   account.Get("lock/<address>/<traceid>", apiv1.Account_lock)
        //   account.Get("status/<address>", apiv1.Account_status)
        // }
        // transaction := test.Group("/transaction")
        // {
        //   transaction.Get("<txhash>",apiv1.Transaction)
        //   transaction.Get("list/<account>/<fromdate>/<todate>", apiv1.Transaction_list)
        //   transaction.Get("lock/<account>/<fromdate>/<todate>", apiv1.Transaction_lock)
        // }
      }

      // v1 := authorized.Group("/v1")
      // {
      // }
	}
	return router
}
