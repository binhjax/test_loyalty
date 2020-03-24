package router

import (
   "github.com/qiangxue/fasthttp-routing"
   "github.com/binhnt-teko/test_loyalty/src/app/api"
	 "github.com/binhnt-teko/test_loyalty/src/app/middleware"
)

func NewRouter()  *routing.Router {
	router := routing.New()

	auth := new(api.ApiAuth)
	apiv1 := new(api.ApiV1)

	api := router.Group("api")
	{
			v1 := api.Group("v1")
			v1.Get("/login", auth.Login_get)
			v1.Post("/login", auth.Login)
			// Authentication required
			authorized := v1.Group("/")

			authorized.Use(middleware.JWTMiddleware())
			{

						cash := authorized.Group("cash")
						{
							cash.Put("credit/<address>/<amount>/<traceid>", apiv1.Cash_credit)
							cash.Put("debit/<address>/<amount>/<traceid>", apiv1.Cash_debit)
							cash.Put("transfer/<from>/<to>/<amount>/<note>/<traceid>", apiv1.Cash_transfer)
						}
						balance := authorized.Group("balance")
						{
							balance.Get("<address>", apiv1.Balance)
							balance.Get("all", apiv1.Balance_all)
						}
						account := authorized.Group("account")
						{
							account.Post("new", apiv1.Account_new)
							account.Get("total", apiv1.Account_total)
							account.Get("list/active", apiv1.Account_list_active)
							account.Get("list/inactive", apiv1.Account_list_inactive)
							account.Get("lock/<address>/<traceid>", apiv1.Account_lock)
							account.Get("status/<address>", apiv1.Account_status)
						}
						transaction := authorized.Group("transaction")
						{
							transaction.Get("<txhash>",apiv1.Transaction)
							transaction.Get("list/<account>/<fromdate>/<todate>", apiv1.Transaction_list)
							transaction.Get("lock/<account>/<fromdate>/<todate>", apiv1.Transaction_lock)
						}
				}
		}
	return router
}
