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
	config.Version = sarama.V0_10_1_0
	config.Producer.Return.Successes = true

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

	fmt.Println("Ready to produce messages. Write something to stdin...")
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		value := stdin.Text()

		msg := &sarama.ProducerMessage{Topic: Topic, Value: sarama.StringEncoder(value)}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Printf("Fail to send message: %s\n", err)
			continue
		}

		fmt.Printf("> Message '%s' sent to partition %d at offset %d\n", value, partition, offset)
	}
}
