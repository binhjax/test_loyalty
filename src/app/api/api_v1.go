package api

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	// "github.com/valyala/fasthttp"
  "github.com/binhnt-teko/test_loyalty/src/app/config"
  "strings"
  // "github.com/go-redis/redis"
  // "encoding/json"
)
type ApiV1 struct {}

func Response(c  *routing.Context, code int, desr string, data interface{} ) error {
		res := &config.ApiResponse{
				Rescode: code,
				Resdecr: desr,
				Resdata: data,
		}
		fmt.Fprintf(c,res.ToJson())
		return nil
}

/***************** Cash credit function  ************/
 func (api ApiV1) Cash_credit(c *routing.Context) error {
     address := c.Param("address")
     amount := c.Param("amount")
     traceid := c.Param("traceid")

     if address == "" {
       fmt.Fprintf(c,"error: Please add from address ")
       return nil
     }
     if amount == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return nil
     }
     if traceid == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return nil
     }

     address = strings.TrimPrefix(address,"0x")
     //
     //
  	 // result, err := utils.TransferToken(from,to,amount,append)
     // if err != nil {
     //       fmt.Fprintf(c,"Error to transfer token: ", err)
     //       return
     // }
		 // fmt.Fprintf(c,"transaction: ", result)
     // fmt.Fprintf(c,"transaction: penđing")
		 return nil
 }

 // call transfer token
 func (api ApiV1) Cash_debit(c *routing.Context) error {
     address := c.Param("address")
     amount := c.Param("amount")
     traceid := c.Param("traceid")

     if address == "" {
       fmt.Fprintf(c,"error: Please add from address ")
       return nil
     }
     if amount == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return nil
     }
     if traceid == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return nil
     }

     address = strings.TrimPrefix(address,"0x")
     // fmt.Fprintf(c,"transaction: penđing")
		  return nil
 }

 // call transfer token
 func (api ApiV1) Cash_transfer(c *routing.Context) error {
     from := c.Param("from")
     to := c.Param("to")
     amount := c.Param("amount")
     note := c.Param("note")
     traceid := c.Param("traceid")

     if from == "" {
       fmt.Fprintf(c,"error: Please add from address ")
       return nil
     }
     if to == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return nil
     }

     from = strings.TrimPrefix(from,"0x")
     to = strings.TrimPrefix(to,"0x")
     //
  	 // result, err := utils.TransferToken(from,to,amount,append)
     // if err != nil {
     //       fmt.Fprintf(c,"Error to transfer token: ", err)
     //       return
     // }
		 // fmt.Fprintf(c,"transaction: ", result)
     fmt.Fprintf(c,"cash_transfer: %v,%v,%v,%v,%v ",from,to,amount,note,traceid)
		 return nil
 }

///// Balance function ///////////
 func (api ApiV1) Balance(c *routing.Context) error {
   address := c.Param("address")

   if address == "" {
     fmt.Fprintf(c,"error: Please add from address ")
     return nil
   }
   address = strings.TrimPrefix(address,"0x")

   if address == "" {
     fmt.Fprintf(c,"error: Please add from address ")
     return nil
   }

    fmt.Fprintf(c,"balance: %v ", address)
		return nil
 }
 // call transfer token
 func (api ApiV1) Balance_all(c *routing.Context) error {
		 fmt.Fprintf(c,"balance_all ")
     // fmt.Fprintf(c,"transaction: penđing")
		 return nil
 }

 //////// Accunt functions //////////
 func (api ApiV1) Account_new(c *routing.Context) error {
   fmt.Println("account_new: start")
	 res := &config.ApiResponse{
			 Rescode: 0,
			 Resdecr: "successfully create new account",
			 Resdata: nil,
	 }
	 fmt.Fprintf(c,res.ToJson())
	 return nil
 }
 func (api ApiV1) Account_total(c *routing.Context) error {
   fmt.Fprintf(c,"account_total: ")
	 return nil
 }

 func (api ApiV1) Account_list_active(c *routing.Context) error  {
    fmt.Fprintf(c,"account_list_active: ")
		return nil
 }

 func (api ApiV1) Account_list_inactive(c *routing.Context) error  {
   fmt.Fprintf(c,"account_list_inactive: ")
	 return nil
 }

 func (api ApiV1) Account_lock(c *routing.Context) error {
    fmt.Fprintf(c,"account_list_inactive: ")
    address := c.Param("address")
    traceid := c.Param("traceid")

    if address == "" {
      fmt.Fprintf(c,"error: Please add from address ")
      return nil
    }
    address = strings.TrimPrefix(address,"0x")
    fmt.Fprintf(c,"account_lock: %v %s %v",address,", traceid: ",traceid)
		return nil
 }
 func (api ApiV1) Account_status(c *routing.Context) error {
     fmt.Fprintf(c,"account_status: ")
     address := c.Param("address")
     if address == "" {
       fmt.Fprintf(c,"error: Please add from address ")
       return nil
     }
     fmt.Fprintf(c,"account_status: ")
		 return nil
 }


 ///// Transaction functions
 func (api ApiV1) Transaction(c *routing.Context) error {
   txhash := c.Param("txhash")

   if txhash == "" {
     fmt.Fprintf(c,"error: Please add txhash ")
     return nil
   }
   txhash = strings.TrimPrefix(txhash,"0x")
   fmt.Fprintf(c,"transaction: %v",txhash)
	 return nil
 }

 func (api ApiV1) Transaction_list(c *routing.Context) error  {
    account := c.Param("account")
    fromdate := c.Param("fromdate")
    todate := c.Param("todate")

   if account == "" {
     fmt.Fprintf(c,"error: Please add account ")
     return nil
   }
   account = strings.TrimPrefix(account,"0x")
   fmt.Fprintf(c,"transaction_list: %v,%v,%v",account,fromdate,todate)
	 return nil
 }
 func (api ApiV1) Transaction_lock(c *routing.Context) error {
   account := c.Param("account")
   fromdate := c.Param("fromdate")
   todate := c.Param("todate")

  if account == "" {
    fmt.Fprintf(c,"error: Please add account ")
    return nil
  }
  account = strings.TrimPrefix(account,"0x")
  fmt.Fprintf(c,"transaction_lock: %v,%v,%v",account,fromdate,todate)
	return nil
 }
