package workers

import (
		// "log"
		// "os"
		// "strings"
  	"encoding/json"
		"github.com/streadway/amqp"
		"fmt"
    "github.com/binhnt-teko/test_loyalty/app/server/config"
)

func NewPubWorker() (*PubWorker, error){
  cfg  := config.Configuration
  url := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.RabbitMQ.User,cfg.RabbitMQ.Password,cfg.RabbitMQ.Host,cfg.RabbitMQ.Port)
  conn, err := amqp.Dial(url)
  if err != nil  {
    fmt.Println("Cannot connect to rabbitmq ", err)
    return nil, err
  }
  //Copy map
  queues := make(map[string]string)
  for k,v := range cfg.RabbitMQ.Queues {
      queues[k] = v
  }
  //Declare quuee for channel
  channel, err := conn.Channel()
  if err != nil  {
    fmt.Println("Cannot get channel ", err)
    return nil, err
  }
  for _,queue_name := range queues {
      fmt.Println("Declare queue: ", queue_name)
       _, err := channel.QueueDeclare(
        queue_name, // name
        true,         // durable
        false,        // delete when unused
        false,        // exclusive
        false,        // no-wait
        nil,          // arguments
      )
      if err != nil {
         fmt.Println("Cannot declare queue: ",queue_name )
      }
  }
  return &PubWorker{
    Conn:  conn,
    Channel: channel,
    Queues: queues,
    Active: true,
  } , nil
}

func (p PubWorker) Publish(queueName string, data interface{} )  error {
	body, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Cannot find ")
		return  err
	}
  // fmt.Println("Publish data: ", string(body))
  //Declare quuee for channel
  // if p.Conn.IsClosed(){
  //   fmt.Println("Cnnection is close ")
  //   return nil
  // }
  //
  return p.Channel.Publish(
		"",     // exchange
		queueName, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
}
func (p PubWorker) Close() error {
	if p.Channel != nil {
			p.Channel.Close()
	}
	if p.Conn != nil {
	 		p.Conn.Close()
	}
	return nil
}
