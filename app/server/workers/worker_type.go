package workers

import (
  // "fmt"
	// "bytes"
	// "log"
	// "time"
  // "math/rand"
  	"github.com/streadway/amqp"
  // "github.com/binhnt-teko/test_loyalty/app/server/config"
)

type PubWorker struct {
  	Conn *amqp.Connection
		Channel *amqp.Channel
		Queues map[string]string
    Active bool
}

type  SubWorker struct {
  	Conn *amqp.Connection
		Channel *amqp.Channel
		Queues map[string]string
		Forever chan bool
    Active bool
}

type Task struct {
   Method string `json:"method"`
   Arguments map[string] interface{} `json:"arguments"`
}

type TaskJob struct {
   Id string `json:"id"`
   Job Task `json:"task"`
   State string `json:"state"`
   Result interface{} `json:"result"`
}

type TaskHandler = func(tj TaskJob,args map[string] interface{})
