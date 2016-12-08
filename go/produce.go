package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

// Produce produces messages from stdin to Kafka
func Produce(cmd *cobra.Command, args []string) {

	var config = sarama.NewConfig()
	config.Net.TLS.Enable = true
	config.Net.SASL.Enable = true
	config.Net.SASL.User = Username
	config.Net.SASL.Password = Password
	config.ClientID = Username
	config.Producer.Return.Successes = true

	var err error
	Producer, err = sarama.NewSyncProducer([]string{BrokerAddr}, config)
	HandleError(err)

	fmt.Println("Ready to produce messages. Write something to stdin...")
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		value := stdin.Text()

		msg := &sarama.ProducerMessage{Topic: Topic, Value: sarama.StringEncoder(value)}

		partition, offset, err := Producer.SendMessage(msg)
		if err != nil {
			fmt.Printf("Fail to send message: %s\n", err)
			continue
		}

		fmt.Printf("> Message '%s' sent to partition %d at offset %d\n", value, partition, offset)
	}
}
