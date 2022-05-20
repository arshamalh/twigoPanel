package main

import (
	"fmt"
	"time"

	"github.com/arshamalh/twigo/entities"
)

var tweet_authors = map[string]string{}
var found_tweets = []string{}

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

func CollectTweetData(tweet entities.Tweet) {
	fmt.Printf("Collecting tweet data: %s\n", tweet.ID)
	item_id := fmt.Sprintf("tweet_%s_written_by_%s", tweet.ID, tweet.AuthorID)
	if _, ok := tweet_authors[tweet.ID]; ok {
		tweet_data := &TweetData{}
		// TODO: Read tweet_data from db
		fmt.Println(item_id)
		fmt.Printf("%#v\n", tweet_data)
		tweet_data.PublicMetrics[time.Now().Unix()] = tweet.PublicMetrics
		// TODO: Write tweet_data on db
	} else {
		tweet_data := &TweetData{
			CreatedAt:     tweet.CreatedAt,
			AuthorID:      tweet.AuthorID,
			ID:            tweet.ID,
			Text:          tweet.Text,
			PublicMetrics: map[int64]entities.TweetPublicMetrics{},
		}
		tweet_data.PublicMetrics[time.Now().Unix()] = tweet.PublicMetrics
		// TODO: Write tweet_data on db
		tweet_authors[tweet.ID] = tweet.AuthorID
	}
}
