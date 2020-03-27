package main

import (
	"fmt"
	"github.com/joho/godotenv"
  "github.com/binhnt-teko/test_loyalty/app/server/account"
  "github.com/binhnt-teko/test_loyalty/app/server/connection"
  "github.com/binhnt-teko/test_loyalty/app/server/config"
	"github.com/binhnt-teko/test_loyalty/app/server/router"
	"github.com/binhnt-teko/test_loyalty/app/server/workers"
	"github.com/binhnt-teko/test_loyalty/app/server/blocklisteners"
	"github.com/binhnt-teko/test_loyalty/app/server/tasks"
  "github.com/binhnt-teko/test_loyalty/app/server/contracts"
	 "github.com/binhnt-teko/test_loyalty/app/server/eventlog"
	"sync"
)

func init() {
	 println("init function")
	 godotenv.Load()
   config.Init()
	 connection.Init()
	 account.Init()
	 workers.Init()
	 listeners.Init()
	 tasks.Init()
	 contracts.Init()
}

func main() {
	 contracts.LContract.Load("../config/contract.yml");

	 var wg sync.WaitGroup
	 wg.Add(4)
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
	 listeners.Start(&wg)
	 go func (){
		 	  println("Start worker 1 to consume message from task queue: ",len(workers.Subs))
			  defer wg.Done()
				workers.Subs[0].Consume("tasks")
	 }()
	 go func (){
			 println("Start worker 1 to consume message from task queue: ",len(workers.Subs))
			 defer wg.Done()
			 workers.Subs[1].Consume("tasks")
	}()
	 go func (){
			println("Loop Api Server ")
			defer wg.Done()
			router.ApiServer()
	 }()
	 go func () {
		 defer wg.Done()
			eventlog.Init()
	 }()

	 wg.Wait()
	 fmt.Println("Finished webserver")
}
