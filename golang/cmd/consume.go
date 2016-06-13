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
		fmt.Printf("Err : %v", err)
		os.Exit(0)
	}
}

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Consume message from the topic and display them.",
	Run: func(cmd *cobra.Command, args []string) {
		var config = sarama.NewConfig()
		config.ClientID = key
		client, err := sarama.NewClient([]string{host}, config)
		HandleError(err)

		offsetManager, err := sarama.NewOffsetManagerFromClient(group, client)
		HandleError(err)

		partitionOffsetManager, err := offsetManager.ManagePartition(topic, 0)
		HandleError(err)

		consumer, err := sarama.NewConsumerFromClient(client)
		HandleError(err)

		offset, _ := partitionOffsetManager.NextOffset()
		partitionConsumer, err := consumer.ConsumePartition(topic, 0, offset)
		HandleError(err)

		defer func() {
			if err := consumer.Close(); err != nil {
				fmt.Printf("Err : %v", err)
			}
			if err := partitionOffsetManager.Close(); err != nil {
				fmt.Printf("Err : %v\n", err)
			}
			if err := partitionConsumer.Close(); err != nil {
				fmt.Printf("Err : %v\n", err)
			}
		}()

		// Trap SIGINT to trigger a shutdown.
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
	ConsumerLoop:
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				fmt.Printf("%s\n", string(msg.Value))
				partitionOffsetManager.MarkOffset(msg.Offset, topic)
			case <-signals:
				break ConsumerLoop
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(consumeCmd)
}
