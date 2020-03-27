package tasks

import (
  "fmt"
  // "os"
	// "bytes"
	// "log"
	// "time"
  // "math/rand"
  // "errors"
  "github.com/binhnt-teko/test_loyalty/app/server/workers"
	// "encoding/json"
)

func Init(){
   workers.SubcribeHandler("Credit",Credit)
}


func Credit(tj workers.TaskJob, args map[string] interface{}) {
    values := make([]interface{}, 0)
    for _, v := range args {
        values = append(values,v)
    }
    fmt.Println("Credit",tj,values)
    tj.State = "Finished"
    tj.Result = "OK"
    workers.UpdateTaskStatus(tj)
}
func Debit(tj workers.TaskJob,args map[string] interface{}) {
        fmt.Println("Credit",tj)
    tj.State = "Finished"
    tj.Result = "OK"

}
func Transfer(tj workers.TaskJob,args map[string] interface{}) {
        fmt.Println("Credit",tj)
    tj.State = "Finished"
    tj.Result = "OK"

}
