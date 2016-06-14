package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produce a message to a given topic",
	Run: func(cmd *cobra.Command, args []string) {
		var config = sarama.NewConfig()
		// Set key as the client id for authentication
		config.ClientID = key

		producer, err := sarama.NewSyncProducer([]string{kafka}, config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer func() {
			if err := producer.Close(); err != nil {
				fmt.Println("Fail to kill producer")
			}
		}()
		s := bufio.NewScanner(os.Stdin)
		// Reads from the scann
		for s.Scan() {
			msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(s.Text())}

			// Send each line to the kafka instance
			partition, offset, err := producer.SendMessage(msg)
			if err != nil {
				fmt.Printf("FAILED to send message: %s\n", err)
			} else {
				fmt.Printf("> %s sent to partition %d at offset %d\n", s.Text(), partition, offset)
			}
		}
	},
}

// init Register the command
func init() {
	RootCmd.AddCommand(produceCmd)
}
