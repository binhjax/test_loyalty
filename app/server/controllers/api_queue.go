package controllers

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	// "github.com/valyala/fasthttp"
  // "github.com/binhnt-teko/test_loyalty/app/server/middleware"
  // "github.com/binhnt-teko/test_loyalty/app/server/helper"
	"github.com/binhnt-teko/test_loyalty/app/server/workers"
  "github.com/binhnt-teko/test_loyalty/app/server/config"
	"time"
	"encoding/json"
)

type ApiQueue struct {}

func (api ApiQueue) Task(c *routing.Context) error {
    var task = workers.Task{}
    err := json.Unmarshal([]byte(c.PostBody()), &task)
    if err != nil {
        fmt.Println(time.Now()," Cannot parse task job ", err)
        return Response(c,99,"Cannot parse task job",err)
    }
    taskId := workers.RandomString(32)
    var taskJob = workers.TaskJob {
      Id: taskId,
      Job: task,
    }
    err = workers.Pub.Publish("tasks",taskJob)
    if err != nil {
      fmt.Println(time.Now()," Cannot submit task to queue ", err)
      return Response(c,99,"Cannot submit task to queue ",err)
    }
		return Response(c,0,"Add task successfully", taskId)
}
func (api ApiQueue) Status(c *routing.Context) error {
    taskId := c.Param("taskid")

		val, err := config.Redis.Get("task."+ taskId).Result()
    if err != nil {
        fmt.Println(time.Now()," Cannot find task: " + taskId, err)
			  return Response(c,99,"Cannot find task: " + taskId,err)
    }
		taskJob := workers.TaskJob{}
		if err := json.Unmarshal([]byte(val), &taskJob); err != nil {
				fmt.Println("")
				fmt.Println(time.Now()," Cannot parse task in taskjob: " + taskId, err)
				return Response(c,99,"Cannot parse task in taskjob: " + taskId,err)
		}
    //Get status and return
    return Response(c,0,"Task result: "+ taskId, taskJob)
}
