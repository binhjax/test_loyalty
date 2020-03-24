package router

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
  "os"
  "github.com/binhnt-teko/test_loyalty/src/app/api"
	 "github.com/binhnt-teko/test_loyalty/src/app/middleware"
	"sync"
)

func NewRouter(){

	router := routing.New()

	apiv1 := new(api.ApiV1)

	api := router.Group("api"){
		v1 := router.Group("v1")
		v1.Get("/login", apiv1.login_get)
		v1.Post("/login", apiv1.login)
		// Authentication required
		authorized := v1.Group("/")
		authorized.Use(middleware.JWTMiddleware()
		{
		// Authentication required
		authorized := v1.Group("/")
		authorized.Use(middleware.JWTMiddleware()
		{
				cash := authorized.Group("cash")
				{
					cash.Put("credit/<address>/<amount>/<traceid>", apiv1.cash_credit)
					cash.Put("debit/<address>/<amount>/<traceid>", apiv1.cash_debit)
					cash.Put("transfer/<from>/<to>/<amount>/<note>/<traceid>", apiv1.cash_transfer)
				}
				balance := authorized.Group("balance")
				{
					balance.Get("<address>", apiv1.balance)
					balance.Get("all", apiv1.balance_all)
				}
				account := authorized.Group("account")
				{
					account.Post("new", apiv1.account_new)
					account.Get("total", apiv1.account_total)
					account.Get("list/active", apiv1.account_list_active)
					account.Get("list/inactive", apiv1.account_list_inactive)
					account.Get("lock/<address>/<traceid>", apiv1.account_lock)
					account.Get("status/<address>", apiv1.account_status)
				}
				transaction := authorized.Group("transaction")
				{
					transaction.Get("<txhash>",apiv1.transaction)
					transaction.Get("list/<account>/<fromdate>/<todate>", apiv1.transaction_list)
					transaction.Get("lock/<account>/<fromdate>/<todate>", apiv1.transaction_lock)
				}
	}


	return router
}
