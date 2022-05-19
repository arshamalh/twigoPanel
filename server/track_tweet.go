package main

import "fmt"

func TrackTweet(tweet_id string) {
	fields := map[string]interface{}{"tweet.fields": []string{"author_id", "created_at", "public_metrics"}}
	response, err := bot.GetTweet(tweet_id, fields)
	if err != nil {
		fmt.Println(err)
	}
	tweet := response.Data
	if tweet.ID != "" {
		CollectTweetData(tweet)
	}
}
