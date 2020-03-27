package workers

import (
	"fmt"
		// "bytes"
		"log"
		// "time"
		"github.com/streadway/amqp"
		"github.com/binhnt-teko/test_loyalty/app/server/config"
)

func NewSubWorker() (*SubWorker,error) {
		cfg  := config.Configuration
		url := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.RabbitMQ.User,cfg.RabbitMQ.Password,cfg.RabbitMQ.Host,cfg.RabbitMQ.Port)
		conn, err := amqp.Dial(url)
		if err != nil  {
			fmt.Println("Cannot connect to rabbitmq ", err)
			return nil, err
		}
		queues := make(map[string]string)
		for k,v := range cfg.RabbitMQ.Queues {
				queues[k] = v
		}
		//Declare quuee for channel
		channel, err := conn.Channel()

		for _,queue_name := range queues {
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
		err = channel.Qos(
			1,     // prefetch count
			0,     // prefetch size
			false, // global
		)
		if err != nil {
			 fmt.Println("Cannot set QoS")
			 return nil, err
		}
		forever := make(chan bool)
	  return &SubWorker{
		  	Conn: conn,
				Channel: channel,
				Queues : queues,
				Forever: forever,
		    Active: true,
		}, nil
}

func (p SubWorker) Consume(queueName string) error {
		 	msgs, err := p.Channel.Consume(
				queueName, // queue
				"",     // consumer
				false,  // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)
			if err != nil  {
				fmt.Println("Cannot consume queue: " + queueName, err)
				return err
			}
			//Start new thread to process delivery channel
			go func() {
				for msg := range msgs {
					// log.Printf("Received a message: %s", string(msg.Body))
					err := ProcessTask(msg.Body)
					if err != nil  {
							//Acknowledge to queue
							fmt.Println("Send error to error queue", err)
					}
					msg.Ack(false)
				}
			}()
			log.Printf(" [*] Waiting for messages. To exit from close")
			<-p.Forever
			return nil
}
func (p SubWorker) Close() error {
	//Stop consume
	p.Forever <- true
	//Stop channel and connections
	if p.Channel != nil {
			p.Channel.Close()
	}
	if p.Conn != nil {
	 		p.Conn.Close()
	}
	return nil
}
