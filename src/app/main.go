package main

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
  "os"
  "github.com/binhnt-teko/test_loyalty/src/app/config"
  "github.com/binhnt-teko/test_loyalty/src/app/redis"
  "github.com/binhnt-teko/test_loyalty/src/app/rpc"
	"sync"
)

var cfg *config.Config
var rpcrouting *utils.RpcRouting
var redisCache *redis.RedisPool

func init() {
  config_file := "./config/config.yaml"
  if len(os.Args) == 2 {
      config_file = os.Args[1]
   }

   println("init function")
   cfg = config.LoadConfig(config_file)
}

func main() {
	//Creat redis connection
	println("Initialize redis")
	redisCache = redis.NewRedisPool()

	println("Delete old data in redis ")
	//utils.DeleteData("transaction*")
	//utils.DeleteData("nonce*")

	 //Load all wallets in hosts
	 println("Create rpc connection pool ")
	 rpcrouting = rpc.NewRouting(cfg)


	 var wg sync.WaitGroup

	 wg.Add(3)

	 go func (){
				println("Loop Routing process message ")
				defer wg.Done()
				rpcrouting.Process()
	 }()

	 go func (){
		 	  println("Loop redisPool ")
			  defer wg.Done()
				redisCache.Process()
	 }()

	 go func (){
			 println("Loop httpServer ")
			defer wg.Done()
			httpServer()
	 }()

	 wg.Wait()
	 fmt.Println("Finished webserver")
}

func httpServer(){
	router := NewRouter()
	server := &fasthttp.Server{
		Name:    "JWT API Server",
		Handler: router.HandleRequest,
	}

	fmt.Println("Start listening")

	if cfg.Webserver.Tls {
		fmt.Println("Start server using TLS ")
		panic(server.ListenAndServeTLS(":"+ cfg.Webserver.Port, cfg.Webserver.CertificateFile,cfg.Webserver.KeyFile))
	} else {
		fmt.Println("Start server without TLS  ")
		panic(server.ListenAndServe(":"+ cfg.Webserver.Port))
	}
}
