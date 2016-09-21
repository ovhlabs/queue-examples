package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

var (
	BrokerAddr    string
	Topic         string
	Username      string
	Password      string
	ConsumerGroup string

	Verbose bool

	Cli *cobra.Command
)

func init() {
	if Verbose {
		sarama.Logger = log.New(os.Stderr, "[sarama] ", log.LstdFlags)
	}

	Cli = &cobra.Command{
		Use:   "kafka-client",
		Short: "Go Kafka client to produce/consume using SASL/SSL",
	}

	Cli.PersistentFlags().StringVar(&BrokerAddr, "broker", "", "Kafka broker address")
	Cli.PersistentFlags().StringVar(&Topic, "topic", "", "Topic")
	Cli.PersistentFlags().StringVar(&Username, "username", "", "SASL username")
	Cli.PersistentFlags().StringVar(&Password, "password", "", "SASL password")
	Cli.PersistentFlags().StringVar(&ConsumerGroup, "consumer-group", "", "Consumer group")
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
	err := Cli.Execute()
	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		fmt.Println("An error occurred: ", err.Error())
		os.Exit(1)
	}
}
