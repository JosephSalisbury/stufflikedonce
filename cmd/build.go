package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/JosephSalisbury/stufflikedonce/client"
	"github.com/JosephSalisbury/stufflikedonce/db"
)

var (
	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Build a local database of liked Tweets",
		Run:   buildRun,
	}
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

func buildRun(cmd *cobra.Command, args []string) {
	log.Printf("building local database of liked Tweets")

	client, err := client.NewTwitterClient(twitterAccessToken, twitterAccessTokenSecret, twitterConsumerKey, twitterConsumerSecret)
	if err != nil {
		log.Fatalf("%v", err)
	}

	database, err := db.NewFileDatabase(fileDatabasePath)
	if err != nil {
		log.Fatalf("%v", err)
	}

	likedTweets, err := client.GetLikedTweets()
	if err != nil {
		log.Fatalf("%v", err)
	}

	for tweet := range likedTweets {
		if err := database.Save(tweet); err != nil {
			log.Fatalf("%v", err)
		}
	}

	log.Printf("built local database of liked Tweets")
}
