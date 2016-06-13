package cmd

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produce a message to a given topic",
	Run: func(cmd *cobra.Command, args []string) {
		var config = sarama.NewConfig()
		config.ClientID = key
		producer, err := sarama.NewSyncProducer([]string{host}, config)
		if err != nil {
			fmt.Printf("Err : %v\n", err)
		}
		defer func() {
			if err := producer.Close(); err != nil {
				fmt.Printf("Err : %v\n", err)
			}
		}()

		msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Printf("FAILED to send message: %s\n", err)
		} else {
			fmt.Printf("> message sent to partition %d at offset %d\n", partition, offset)
		}
	},
}

func init() {
	RootCmd.AddCommand(produceCmd)
}
