package controllers

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
  "github.com/binhnt-teko/test_loyalty/app/server/config"
  // "strings"
)

type FilePath struct {
   Filepath string `json:"filepath"`
}

type JwtFile struct {
  Filepath string `json:"filepath"`
}

type FilterLog struct {
  FromBlock int64 `json:"FromBlock"`
	ToBlock int64 `json:"ToBlock"`
	Addresses []string `json:"Addresses"`
}

func Response(c  *routing.Context, code int, desr string, data interface{} ) error {
		res := &config.ApiResponse{
				Rescode: code,
				Resdecr: desr,
				Resdata: data,
		}
		fmt.Fprintf(c,res.ToJson())
		return nil
}
