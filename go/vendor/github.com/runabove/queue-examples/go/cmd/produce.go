package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produce messages",
	Run: func(cmd *cobra.Command, args []string) {

		var config = sarama.NewConfig()
		config.ClientID = Username
		config.Net.TLS.Enable = true
		config.Net.SASL.Enable = true
		config.Net.SASL.User = Username
		config.Net.SASL.Password = Password

		fmt.Println(BrokerAddr)
		fmt.Println(config)

		producer, err := sarama.NewSyncProducer([]string{BrokerAddr}, config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer func() {
			if err := producer.Close(); err != nil {
				fmt.Println("Fail to close producer")
			}
		}()

		fmt.Println("Write to STDIN to send a message:")
		s := bufio.NewScanner(os.Stdin)
		// Read from STDIN
		for s.Scan() {
			msg := &sarama.ProducerMessage{Topic: Topic, Value: sarama.StringEncoder(s.Text())}

			partition, offset, err := producer.SendMessage(msg)
			if err != nil {
				fmt.Printf("Fail to send message: %s\n", err)
				continue
			}

			fmt.Printf("> %s sent to partition %d at offset %d\n", s.Text(), partition, offset)
		}
	},
}

func init() {
	RootCmd.AddCommand(produceCmd)
}
