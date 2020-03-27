package controllers

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/binhnt-teko/test_loyalty/app/server/config"
	"time"
	"encoding/json"
)

type ApiAccount struct {}

//Load account from keystore
func (api ApiAccount) LoadKeystore(c  *routing.Context) error {
    var keystore = &FilePath{}
  	err := json.Unmarshal([]byte(c.PostBody()), keystore)
  	if err != nil {
  			fmt.Println(time.Now()," Cannot parse keystore path ", err)
        return Response(c,99,"Cannot parse keystore path",err)
  	}

    err = config.LoadAccountsFromKeyStore(keystore.Filepath)
    if err != nil {
      fmt.Println(time.Now()," Cannot load keystore to redis ", err)
      return Response(c,99,"Cannot load keystore to redis",err)
    }
    return Response(c,0,"LoadKeystore successfully",nil)
}
//Load account from file
func (api ApiAccount) LoadFile(c  *routing.Context) error {
    var keystore = &FilePath{}
  	err := json.Unmarshal([]byte(c.PostBody()), keystore)
  	if err != nil {
  			fmt.Println(time.Now()," Cannot parse keystore path ", err)
        return Response(c,99,"Cannot parse keystore path",err)
  	}

    err = config.LoadAccountsFromFile(keystore.Filepath)
    if err != nil {
      fmt.Println(time.Now()," Cannot load keystore to redis ", err)
      return Response(c,99,"Cannot load keystore to redis",err)
    }
    return Response(c,0,"LoadKeystore successfully",nil)
}
//List account in redis
func (api ApiAccount) ListToken(c  *routing.Context) error {
    tokens, err  := config.Redis.Keys("token.*").Result()
    if err != nil {
      // handle error
      fmt.Println(" Cannot get addresses ")
      return Response(c, 99 ,"Cannot get addresses",err)
    }
    accounts := []config.TokenAccount{}
    for _, token := range tokens {
      val, err := config.Redis.Get(token).Result()
      if err != nil {
          fmt.Println(time.Now()," Cannot find token: ", token)
          continue
      }
      account := config.TokenAccount{}
      err = json.Unmarshal([]byte(val), &account)
      if err != nil {
          fmt.Println(time.Now()," Cannot parse account ", err)
          continue
      }
      accounts = append(accounts,account)
    }
    return Response(c,0,"List Accounts",accounts)
}
func (api ApiAccount) SaveFile(c  *routing.Context) error {
    var keystore = &FilePath{}
    err := json.Unmarshal([]byte(c.PostBody()), keystore)
    if err != nil {
        fmt.Println(time.Now()," Cannot parse keystore path ", err)
        return Response(c,99,"Cannot parse keystore path",err)
    }

    err = config.SaveAccountsToFile(keystore.Filepath)
    if err != nil {
      fmt.Println(time.Now()," Cannot save account from redis to file ", err)
      return Response(c,99,"Cannot save account from redis to file",err)
    }
    return Response(c,0,"SaveToken successfully",nil)
}
