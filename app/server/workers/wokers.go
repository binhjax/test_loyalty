package workers

import (
  "fmt"
  // "os"
	// "bytes"
	// "log"
	// "time"
  "math/rand"
  // "github.com/binhnt-teko/test_loyalty/app/server/config"
	// "encoding/json"
)
var Pub *PubWorker
var Subs []*SubWorker

func Init() error {
  fmt.Println("Initialize workers")
  pub, err := NewPubWorker()
  if err != nil {
    fmt.Println("Init Pubworker failed")
    return err
  }
  Pub = pub
  subs := make([]*SubWorker,0)
  for i := 0; i < 2; i++ {
    sub, err := NewSubWorker()
    if err != nil {
      fmt.Println("Init SubWorker failed")
      continue
    }
    subs  = append(subs,sub)
	}
  if len(subs) == 0 {
      fmt.Println("Cannot init subworker => no worker")
  }
  fmt.Println("Total workers:", len(subs))
  Subs = subs
  // Sub = &SubWorker{}
  Handlers = make(map[string]TaskHandler,0)

  return nil
}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
