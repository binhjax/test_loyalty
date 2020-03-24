package router

import (
	"fmt"
	// "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
  "github.com/binhnt-teko/test_loyalty/src/app/config"
  // "os"
  // "github.com/binhnt-teko/test_loyalty/src/app/api"
	 // "github.com/binhnt-teko/test_loyalty/src/app/middleware"
	// "sync"
)

func ApiServer(){
	router := NewRouter()
	server := &fasthttp.Server{
		Name:    "JWT API Server",
		Handler: router.HandleRequest,
	}

	fmt.Println("Start listening")

	if config.Configuration.Webserver.Tls {
		fmt.Println("Start server using TLS ")
		panic(server.ListenAndServeTLS(":"+ config.Configuration.Webserver.Port, config.Configuration.Webserver.CertificateFile,config.Configuration.Webserver.KeyFile))
	} else {
		fmt.Println("Start server without TLS  ")
		panic(server.ListenAndServe(":"+ config.Configuration.Webserver.Port))
	}
}
