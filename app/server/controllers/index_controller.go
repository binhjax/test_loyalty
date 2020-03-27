package controllers

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
  // "github.com/binhnt-teko/test_loyalty/app/server/config"
  // "strings"
)
type IndexController struct {}

func (api *IndexController) Index(c  *routing.Context) error {
		fmt.Println("API Access Login")
		return Response(c,99,"Please use POST method",nil)
}
