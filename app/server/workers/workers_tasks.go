package workers

import (
  "fmt"
  // "os"
	// "bytes"
	// "log"
	// "time"
  // "math/rand"
  "errors"
  "github.com/binhnt-teko/test_loyalty/app/server/config"
	"encoding/json"
)

var Handlers map[string]TaskHandler

func SubcribeHandler(name string, handler TaskHandler){
   Handlers[name] = handler
}

func ProcessTask(data []byte) error {
  taskJob := TaskJob{}
  if err := json.Unmarshal([]byte(data), &taskJob); err != nil {
	   	fmt.Println("Cannot parse task in taskjob")
      //Send to error queue
	}
  //Store message to redis queue
  taskJob.State = "Processing"
  if err := UpdateTaskStatus(taskJob); err != nil {
    	fmt.Println("Cannot update task status")
      return err
  }

  if handler, ok := Handlers[taskJob.Job.Method]; ok {
      fmt.Println("call handler ")
      handler(taskJob,taskJob.Job.Arguments)
  } else {
      // fmt.Printf("Cannot process %s.\n", taskJob.Job.Method)
      return errors.New("Cannot process method")
  }
  return nil
}
func UpdateTaskStatus(taskJob TaskJob) error{
    value, err := json.Marshal(taskJob)
    if err != nil {
        fmt.Println("Cannot convert to json: ", err)
        //Send to error queue
        return err
    }
    err = config.Redis.Set("task." + taskJob.Id,string(value), 0).Err()
    if err != nil {
      fmt.Println("Cannot save to redis : ", err)
      //Send to error queue
      return err
    }
    return nil
}
