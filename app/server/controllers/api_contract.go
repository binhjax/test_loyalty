package controllers

import (
		"fmt"
		"github.com/qiangxue/fasthttp-routing"
		// "github.com/valyala/fasthttp"
	  "github.com/binhnt-teko/test_loyalty/app/server/contracts"
	  // "github.com/binhnt-teko/test_loyalty/app/server/config"
		// "os"
	  // "io/ioutil"
		// "gopkg.in/yaml.v2"
		"time"
		"encoding/json"
)

type ApiContract struct {}

func (api ApiContract) Load(c *routing.Context) error {
		var filepath = &FilePath{}
		err := json.Unmarshal([]byte(c.PostBody()), filepath)
		if err != nil {
				fmt.Println(time.Now()," Cannot parse filepath path ", err)
				return Response(c,99,"Cannot parse filepath path",err)
		}

	  if err := contracts.LContract.Load(filepath.Filepath); err != nil {
				return Response(c,98,"Load contract error ",err)
		}
		return Response(c,0,"Load contract ", contracts.LContract)
}
func (api ApiContract) View(c *routing.Context) error {
		return Response(c,0,"Load contract ", contracts.LContract)
}
func (api ApiContract) Deploy(c *routing.Context) error {
		if addr,tx,  err := contracts.LContract.Deploy(); err != nil {
				return Response(c,98,"Deploy contract error ",err)
		} else {
				return Response(c,0,"Deploy contract successfully: " + tx, addr)
		}

}

func (api ApiContract) Save(c *routing.Context) error {
		var filepath = &FilePath{}
		err := json.Unmarshal([]byte(c.PostBody()), filepath)
		if err != nil {
				fmt.Println(time.Now()," Cannot parse filepath path ", err)
				return Response(c,99,"Cannot parse filepath path",err)
		}

		if err := contracts.LContract.Save(filepath.Filepath); err != nil {
				return Response(c,98,"Save contract error ",err)
		}
		return Response(c,0,"Save contract successfully", nil)
}
func (api ApiContract) Code(c *routing.Context) error {
		code, err := contracts.LContract.CodeAt()
		if err != nil {
				return Response(c,98,"Get code of contract error ",err)
		}
		return Response(c,0,"Contract Code", code)
}
