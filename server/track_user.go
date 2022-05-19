package main

import (
	"fmt"
	"time"
)

func TrackUser(user_id string) {
	start_time := time.Now().UTC().Add(-5 * time.Minute)
	params := map[string]interface{}{"max_results": 5, "start_time": start_time, "tweet.fields": []string{"created_at"}}
	user_tweets, _ := bot.GetUserTweets(user_id, params)
	if len(user_tweets.Data) != 0 {
		tweet_id := user_tweets.Data[0].ID
		if !contains(found_tweets, tweet_id) {
			fmt.Printf("A tweet found: https://twitter.com/i/web/status/%s\n", tweet_id)
			scheduler.Every(5).Seconds().Tag(fmt.Sprintf("tweet_track_%s", tweet_id)).Do(TrackTweet, tweet_id)
			found_tweets = append(found_tweets, tweet_id)
		}
	}
}
