package cmd

import (
	"log"

	"github.com/JosephSalisbury/stufflikedonce/client"
	"github.com/JosephSalisbury/stufflikedonce/db"
	"github.com/spf13/cobra"
)

var (
	postCmd = &cobra.Command{
		Use:   "post",
		Short: "Post a liked Tweet",
		Run:   postRun,
	}
)

func init() {
	rootCmd.AddCommand(postCmd)
}

func postRun(cmd *cobra.Command, args []string) {
	log.Printf("posting a liked Tweet")

	client, err := client.NewTwitterClient(twitterAccessToken, twitterAccessTokenSecret, twitterConsumerKey, twitterConsumerSecret)
	if err != nil {
		log.Fatalf("%v", err)
	}

	database, err := db.NewFileDatabase(fileDatabasePath)
	if err != nil {
		log.Fatalf("%v", err)
	}

	randomLikedTweet, err := database.GetRandomLikedTweet()
	if err != nil {
		log.Fatalf("%v", err)
	}

	if err := client.Post(randomLikedTweet); err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("posted a liked Tweet")
}
