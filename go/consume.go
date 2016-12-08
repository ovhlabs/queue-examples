package main

import (
	"fmt"

	"github.com/Shopify/sarama"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/bsm/sarama-cluster.v2"
)

// Consume consumes messages from Kafka
func Consume(cmd *cobra.Command, args []string) {

	var config = sarama.NewConfig()
	config.Net.TLS.Enable = true
	config.Net.SASL.Enable = true
	config.Net.SASL.User = Username
	config.Net.SASL.Password = Password
	config.Version = sarama.V0_10_0_1
	config.ClientID = Username

	if ConsumerGroup == "" {
		ConsumerGroup = Username + ".go"
	}

	clusterConfig := cluster.NewConfig()
	clusterConfig.Config = *config
	clusterConfig.Consumer.Return.Errors = true
	clusterConfig.Group.Return.Notifications = true

	var err error
	Consumer, err = cluster.NewConsumer([]string{BrokerAddr}, ConsumerGroup, []string{Topic}, clusterConfig)
	HandleError(err)

	// Consume errors
	go func() {
		for err := range Consumer.Errors() {
			log.WithError(err).Error("Error during consumption")
		}
	}()

	// Consume notifications
	go func() {
		for note := range Consumer.Notifications() {
			log.WithField("note", note).Debug("Rebalanced consumer")
		}
	}()

	fmt.Println("Ready to consume messages...")
	// Consume messages
	for msg := range Consumer.Messages() {

		log.WithFields(log.Fields{
			"consumerGroup": ConsumerGroup,
			"topic":         Topic,
			"timestamp":     msg.Timestamp,
			"key":           string(msg.Key),
			"value":         string(msg.Value),
			"partition":     msg.Partition,
			"offset":        msg.Offset,
		}).Debug("Consume message")

		Consumer.MarkOffset(msg, "")

		fmt.Println(string(msg.Value))
	}
}
