package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var host string
var group string
var key string
var topic string
var message string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "qaas-client",
	Short: "Qaas client using authentication to produce/consume on a topic",
	Long:  "",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&topic, "topic", "", "Destination topic")
	RootCmd.PersistentFlags().StringVar(&group, "group", "", "The group to use as a consumer")
	RootCmd.PersistentFlags().StringVar(&host, "host", "", "The host to connect to")
	RootCmd.PersistentFlags().StringVar(&key, "key", "", "The authentication key to use")
	RootCmd.PersistentFlags().StringVar(&message, "message", "my grumpy message", "The message to send")
}
