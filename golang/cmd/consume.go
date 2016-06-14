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

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Consume message from the topic and display them.",
	Run: func(cmd *cobra.Command, args []string) {

		//Create a new client
		var config = sarama.NewConfig()
		config.ClientID = key
		client, err := sarama.NewClient([]string{host}, config)
		HandleError(err)

		//create an offsetManager
		offsetManager, err := sarama.NewOffsetManagerFromClient(group, client)
		HandleError(err)

		//create a client
		consumer, err := sarama.NewConsumerFromClient(client)
		HandleError(err)

		//create the message chan, that will receive the queue
		messagesChan := make(chan string)

		// read the number of partition for the given topic
		partitions, err := consumer.Partitions(topic)
		HandleError(err)

		//create a consumer for each partition
		for _, p := range partitions {
			partitionOffsetManager, err := offsetManager.ManagePartition(topic, p)
			HandleError(err)
			defer partitionOffsetManager.AsyncClose()

			// Start a consumer at next offset
			offset, _ := partitionOffsetManager.NextOffset()
			partitionConsumer, err := consumer.ConsumePartition(topic, p, offset)
			HandleError(err)
			defer partitionConsumer.AsyncClose()

			//asynchronously handle message
			go consumptionHandler(partitionConsumer, partitionOffsetManager, messagesChan)
		}

		// Trap SIGINT to trigger a shutdown.
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
	ConsumerLoop:
		for {
			select {
			case msg := <-messagesChan:
				fmt.Printf("%s\n", msg)
			case <-signals:
				break ConsumerLoop
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(consumeCmd)
}

// consumptionHandler pipes the handled messages and push them to a chan
func consumptionHandler(pc sarama.PartitionConsumer, po sarama.PartitionOffsetManager, messagesChan chan string) {
	for {
		select {
		case msg := <-pc.Messages():
			// Write message consumed in the sub channel
			messagesChan <- string(msg.Value)
			po.MarkOffset(msg.Offset, topic)
		case err := <-pc.Errors():
			fmt.Println(err)
		case offsetErr := <-po.Errors():
			fmt.Println(offsetErr)
		}
	}
}
