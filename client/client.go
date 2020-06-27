package client

import (
	"errors"
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Twitter struct {
	client *twitter.Client
}

func NewTwitterClient(accessToken string, accessTokenSecret string, consumerKey string, consumerSecret string) (*Twitter, error) {
	if accessToken == "" {
		return nil, errors.New("twitter access token cannot be empty")
	}
	if accessTokenSecret == "" {
		return nil, errors.New("twitter access token secret cannot be empty")
	}
	if consumerKey == "" {
		return nil, errors.New("twitter consumer key cannot be empty")
	}
	if consumerSecret == "" {
		return nil, errors.New("twitter consumer secret cannot be empty")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	t := &Twitter{
		client: client,
	}

	return t, nil
}

func (t *Twitter) GetLikedTweets() (chan twitter.Tweet, error) {
	c := make(chan twitter.Tweet)

	go func(chan twitter.Tweet) {
		keep_going := true
		var lastTweet twitter.Tweet

		for keep_going {
			likedTweets, _, err := t.client.Favorites.List(&twitter.FavoriteListParams{
				ScreenName: "salisbury_joe",
				Count:      200,
				MaxID:      lastTweet.ID,
			})
			if err != nil {
				log.Fatalf("%v", err)
			}

			if len(likedTweets) < 100 {
				keep_going = false
			}

			for _, tweet := range likedTweets {
				if tweet.ID == lastTweet.ID {
					continue
				}

				c <- tweet
				lastTweet = tweet
			}

			time.Sleep(20 * time.Second)
		}

		close(c)
	}(c)

	return c, nil
}

func (t *Twitter) Post(url string) error {
	if _, _, err := t.client.Statuses.Update(url, &twitter.StatusUpdateParams{}); err != nil {
		return err
	}

	return nil
}
