package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
	"gopkg.in/bsm/sarama-cluster.v2"
)

var (
	BrokerAddr    string
	Topic         string
	Username      string
	Password      string
	ConsumerGroup string

	Verbose bool

	Cli *cobra.Command

	Producer sarama.SyncProducer
	Consumer *cluster.Consumer
)

func init() {
	//sarama.Logger = log.New(os.Stderr, "[debug-sarama] ", log.LstdFlags)

	Cli = &cobra.Command{
		Use:   "kafka-client",
		Short: "Go Kafka client to produce/consume using SASL/SSL",
	}

	Cli.PersistentFlags().StringVar(&BrokerAddr, "broker", "", "Kafka broker address")
	Cli.PersistentFlags().StringVar(&Topic, "topic", "", "Topic, prefixed by your namespace (eg. --topic=myns.topic)")
	Cli.PersistentFlags().StringVar(&Username, "username", "", "SASL username, prefixed by your namespace (eg. --username=myns.user)")
	Cli.PersistentFlags().StringVar(&Password, "password", "", "SASL password")
	Cli.PersistentFlags().StringVar(&ConsumerGroup, "consumer-group", "", "Consumer group, prefixed by your SASL username (eg. --consumer-group=myns.user.g1)")
	Cli.PersistentFlags().BoolVar(&Verbose, "verbose", false, "Verbose mode")

	Cli.AddCommand(&cobra.Command{
		Use:   "consume",
		Short: "Consume messages",
		Run:   Consume,
	})

	Cli.AddCommand(&cobra.Command{
		Use:   "produce",
		Short: "Produce messages",
		Run:   Produce,
	})
}

func main() {
	SetupGracefulShutdown()

	err := Cli.Execute()
	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		fmt.Println("An error occurred: ", err.Error())
		os.Exit(1)
	}
}

func SetupGracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if Producer != nil {
			Producer.Close()
		}
		if Consumer != nil {
			Consumer.Close()
		}
		os.Exit(1)
	}()
}
