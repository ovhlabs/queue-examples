package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

// Consume consumes messages from Kafka
func Consume(cmd *cobra.Command, args []string) {
	var config = sarama.NewConfig()
	config.Net.TLS.Enable = true
	config.Net.SASL.Enable = true
	config.Net.SASL.User = Username
	config.Net.SASL.Password = Password
	config.ClientID = Username

	client, err := sarama.NewClient([]string{BrokerAddr}, config)
	HandleError(err)

	if ConsumerGroup == "" {
		ConsumerGroup = Username + ".go"
	}

	// Create an offsetManager
	offsetManager, err := sarama.NewOffsetManagerFromClient(ConsumerGroup, client)
	HandleError(err)

	// Create a consumer
	consumer, err := sarama.NewConsumerFromClient(client)
	HandleError(err)

	// Create the message chan, that will receive the messages
	messagesChan := make(chan string)

	// read the number of partition for the given topic
	partitions, err := consumer.Partitions(Topic)
	HandleError(err)

	// Create a consumer for each partition
	for _, p := range partitions {
		partitionOffsetManager, err := offsetManager.ManagePartition(Topic, p)
		HandleError(err)
		defer partitionOffsetManager.AsyncClose()

		// Start a partition consumer from the next offset
		offset, _ := partitionOffsetManager.NextOffset()
		partitionConsumer, err := consumer.ConsumePartition(Topic, p, offset)
		HandleError(err)
		defer partitionConsumer.AsyncClose()

		// Asynchronously consume messages
		go consumptionHandler(partitionConsumer, partitionOffsetManager, messagesChan)
	}

	fmt.Println("Ready to consume messages...")

	// Trap SIGINT to trigger a shutdown
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

consumerLoop:
	for {
		select {
		case msg := <-messagesChan:
			fmt.Printf("%s\n", msg)
		case <-sigint:
			close(messagesChan)
			consumer.Close()
			client.Close()
			break consumerLoop
		}
	}

}

// ConsumptionHandler pipes the consumed messages and push them to a chan
func consumptionHandler(pc sarama.PartitionConsumer, pom sarama.PartitionOffsetManager, messagesChan chan string) {
	for {
		select {
		case msg := <-pc.Messages():
			messagesChan <- string(msg.Value)
			pom.MarkOffset(msg.Offset+1, Topic)

		case err := <-pc.Errors():
			fmt.Println(err)

		case offsetErr := <-pom.Errors():
			fmt.Println(offsetErr)
		}
	}
}
