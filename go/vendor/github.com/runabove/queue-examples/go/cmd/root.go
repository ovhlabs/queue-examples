package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

var (
	BrokerAddr string
	Topic      string
	Username   string
	Password   string

	Verbose bool
)

var RootCmd = &cobra.Command{
	Use:   "Go Kafka client example",
	Short: "Go Kafka client using SASL/SSL to produce/consume",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&BrokerAddr, "broker", "", "Kafka broker address")
	RootCmd.PersistentFlags().StringVar(&Topic, "topic", "", "Topic")
	RootCmd.PersistentFlags().StringVar(&Username, "username", "", "SASL username")
	RootCmd.PersistentFlags().StringVar(&Password, "password", "", "SASL password")
	RootCmd.PersistentFlags().BoolVar(&Verbose, "verbose", true, "Verbose mode")

	if Verbose {
		sarama.Logger = log.New(os.Stderr, "[sarama] ", log.LstdFlags)
	}
}
