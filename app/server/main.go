package main

import (
	"fmt"
	"github.com/joho/godotenv"
  "github.com/binhnt-teko/test_loyalty/app/server/config"
	"github.com/binhnt-teko/test_loyalty/app/server/router"
	"sync"
)

func init() {
	 println("init function")
	 godotenv.Load()
   config.Init()
}

func main() {
	 var wg sync.WaitGroup
	 wg.Add(1)
	 // go func (){
		// 		println("Loop Routing process message ")
		// 		defer wg.Done()
		// 		rpcrouting.Process()
	 // }()
	 //
	 // go func (){
		//  	  println("Loop redisPool ")
		// 	  defer wg.Done()
		// 		redisCache.Process()
	 // }()

	 go func (){
			println("Loop Api Server ")
			defer wg.Done()
			router.ApiServer()
	 }()

	 wg.Wait()
	 fmt.Println("Finished webserver")
}
