/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bigraptor",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bigraptor.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".bigraptor" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".bigraptor")
	}

	viper.AutomaticEnv() // read in environment variables that match

	configFileName := ".bigraptor.yaml"

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}

	writeDefaultConfig(configFileName)

	// Access the configuration values
	region := viper.GetString("sns.region")
	topicArn := viper.GetString("sns.topic-arn")
	// ... access other configuration values as needed ...
	fmt.Printf("sns.region: %s\n", region)
	fmt.Printf("sns.topic-arn: %s\n", topicArn)
}

func writeDefaultConfig(configFileName string) {
	yamlExample := []byte(`
sns:
  region: us-west-2
  topic-arn: arn:aws:sns:us-west-2:123456789012:example-topic
  apple: pear
sqs:
  queue-arn: arn:aws:sqs:us-west-2:123456789012
  queue-url: https://sqs.us-west-2.amazonaws.com/123456789012/somename
  region: us-west-2
s3bucket:
  name: mybucket
  region: us-west-2
  s3path: .bigraptor.yaml
client:
  push-frequency: 4m
`)

	// Initialize the default configuration
	err := viper.MergeConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		fmt.Println("Error initializing default config:", err)
		os.Exit(1)
	}

	// Save the merged configuration, including new default values
	err = viper.WriteConfig()
	if err != nil {
		fmt.Println("Error writing config file:", err)
		os.Exit(1)
	}

	fmt.Printf("Default configuration written to %s\n", configFileName)
}
