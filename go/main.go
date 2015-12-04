package main

import (
  "fmt"
  "github.com/codegangsta/cli"
  "os"
  "os/signal"
	"github.com/Shopify/sarama"
)


func main(){
	//parse option
	app := cli.NewApp()
	app.Name = "cons_prod_tester"
	app.Usage = "Consumer producer tester for go"
	app.Action = acceptor

	app.Flags = []cli.Flag{
    cli.StringFlag{
      Name:"kafka, k",
      Value: "127.0.01:9092",
    },
    cli.StringFlag{
      Name:"auth, a",
      Value: "",
    },
    cli.StringFlag{
      Name:"topic, t",
      Value: "topic1",
    },
    cli.StringFlag{
      Name:"group, g",
      Value: "qaas-client",
    },
    cli.StringFlag{
      Name:"mode, m",
      Value: "prod",
    },
	}

	app.Run(os.Args)
}
func acceptor(c *cli.Context){
  k := c.String("kafka")
  m := c.String("mode")
  a := c.String("auth")
  g := c.String("group")
  t := c.String("topic")

  var config = sarama.NewConfig()
  config.ClientID = a
  if m == "cons" {
    err := consumer(k, g, t, config)
    if err != nil {
      fmt.Printf("Err : %v\n", err)
      panic(err)
    }
  } else {
    producer(k, g , t, config)
  }
}
func consumer(url, group, topic string, config *sarama.Config) error {
  client, err := sarama.NewClient([]string{url}, config)
	if err != nil {
		return  err
	}
	offsetManager, err := sarama.NewOffsetManagerFromClient(group, client)
	if err != nil {
		return  err
	}
	partitionOffsetManager, err := offsetManager.ManagePartition(topic, 0)
	if err != nil {
		return  err
	}
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return  err
	}
  defer func() {
      if err := consumer.Close(); err != nil {
        fmt.Printf("Err : %v", err)
      }
  }()
	offset, _ := partitionOffsetManager.NextOffset()
  fmt.Printf("%d\n", offset)
	if err != nil {
		return  err
	}
	partitionConsumer, err := consumer.ConsumePartition(topic, 0,offset)
  if err != nil {
    return err
  }

  defer func() {
      if err := partitionConsumer.Close(); err != nil {
        fmt.Printf("Err : %v\n", err)
      }
  }()

  // Trap SIGINT to trigger a shutdown.
  signals := make(chan os.Signal, 1)
  signal.Notify(signals, os.Interrupt)

  consumed := 0
  ConsumerLoop:
  for {
      select {
      case msg := <-partitionConsumer.Messages():
          fmt.Printf("Data : %v\n", msg)
				  partitionOffsetManager.MarkOffset(msg.Offset, topic)
          consumed++
      case <-signals:
          break ConsumerLoop
      }
  }

  fmt.Printf("Consumer %d\n", consumed)
  return nil
}

func producer(url, group, t string, config *sarama.Config){
  fmt.Println("Producer")
  producer, err := sarama.NewSyncProducer([]string{url}, config)
  if err != nil {
      fmt.Printf("Err : %v\n", err)
  }
  defer func() {
      if err := producer.Close(); err != nil {
        fmt.Printf("Err : %v\n", err)
      }
  }()

  msg := &sarama.ProducerMessage{Topic: "topic1", Value: sarama.StringEncoder("testing 123")}
  partition, offset, err := producer.SendMessage(msg)
  if err != nil {
      fmt.Printf("FAILED to send message: %s\n", err)
  } else {
      fmt.Printf("> message sent to partition %d at offset %d\n", partition, offset)
  }
}
