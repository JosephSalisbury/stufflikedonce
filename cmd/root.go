package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "stufflikedonce",
	}

	twitterAccessToken       string
	twitterAccessTokenSecret string
	twitterConsumerKey       string
	twitterConsumerSecret    string

	fileDatabasePath string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&twitterAccessToken, "twitter-access-token", "", "access token for Twitter")
	rootCmd.PersistentFlags().StringVar(&twitterAccessTokenSecret, "twitter-access-token-secret", "", "access token secret for Twitter")
	rootCmd.PersistentFlags().StringVar(&twitterConsumerKey, "twitter-consumer-key", "", "consumer key for Twitter")
	rootCmd.PersistentFlags().StringVar(&twitterConsumerSecret, "twitter-consumer-secret", "", "consumer secret for Twitter")

	rootCmd.PersistentFlags().StringVar(&fileDatabasePath, "file-database-path", "", "path to file database")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("%v", err)
	}
}
