package router

import (
	"fmt"
	"github.com/valyala/fasthttp"
  "github.com/binhnt-teko/test_loyalty/app/server/config"
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
		fmt.Println("Start server without TLS : " + config.Configuration.Webserver.Port)
		panic(server.ListenAndServe(":"+ config.Configuration.Webserver.Port))
	}
}
