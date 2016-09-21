package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Consume messages",
	Run: func(cmd *cobra.Command, args []string) {

		var config = sarama.NewConfig()
		config.Net.TLS.Enable = true
		config.Net.SASL.Enable = true
		config.Net.SASL.User = Username
		config.Net.SASL.Password = Password

		client, err := sarama.NewClient([]string{BrokerAddr}, config)
		HandleError(err)

		// Create an offsetManager
		offsetManager, err := sarama.NewOffsetManagerFromClient(Username+".group1", client)
		HandleError(err)

		// Create a client
		consumer, err := sarama.NewConsumerFromClient(client)
		HandleError(err)

		// Create the message chan, that will receive the queue
		messagesChan := make(chan string)

		// read the number of partition for the given topic
		partitions, err := consumer.Partitions(Topic)
		HandleError(err)

		// Create a consumer for each partition
		for _, p := range partitions {
			partitionOffsetManager, err := offsetManager.ManagePartition(Topic, p)
			HandleError(err)
			defer partitionOffsetManager.AsyncClose()

			// Start a consumer at next offset
			offset, _ := partitionOffsetManager.NextOffset()
			partitionConsumer, err := consumer.ConsumePartition(Topic, p, offset)
			HandleError(err)
			defer partitionConsumer.AsyncClose()

			// Asynchronously handle message
			go consumptionHandler(partitionConsumer, partitionOffsetManager, messagesChan)
		}

		// Trap SIGINT to trigger a shutdown
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)

	consumerLoop:
		for {
			select {
			case msg := <-messagesChan:
				fmt.Printf("%s\n", msg)
			case <-sigint:
				break consumerLoop
			}
		}
	},
}

// init Register the command
func init() {
	RootCmd.AddCommand(consumeCmd)
}

// ConsumptionHandler pipes the handled messages and push them to a chan
func consumptionHandler(pc sarama.PartitionConsumer, po sarama.PartitionOffsetManager, messagesChan chan string) {
	for {
		select {
		case msg := <-pc.Messages():
			// Write message consumed in the sub channel
			messagesChan <- string(msg.Value)
			po.MarkOffset(msg.Offset+1, Topic)

		case err := <-pc.Errors():
			fmt.Println(err)
		case offsetErr := <-po.Errors():
			fmt.Println(offsetErr)
		}
	}
}
