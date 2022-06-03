package main

import (
	"fmt"
	"time"

	"github.com/arshamalh/twigo/entities"
	"github.com/arshamalh/twigoPanel/database"
)

var tracking_tweets = []string{}

func TrackUser(user_id string) {
	start_time := time.Now().UTC().Add(-5 * time.Minute)
	params := map[string]interface{}{"max_results": 5, "start_time": start_time, "tweet.fields": []string{"created_at"}}
	user_tweets, _ := bot.GetUserTweets(user_id, params)
	if len(user_tweets.Data) != 0 {
		for _, tweet := range user_tweets.Data {
			if !contains(tracking_tweets, tweet.ID) {
				fmt.Printf("A tweet found: https://twitter.com/i/web/status/%s\n", tweet.ID)
				job_tag := fmt.Sprintf("tweet_track_%s", tweet.ID)
				scheduler.Every(5).Seconds().Tag(job_tag).Do(TrackTweet, tweet.ID)
				tracking_tweets = append(tracking_tweets, tweet.ID)
			}
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
		fmt.Printf("Collecting tweet data: %s\n", tweet.ID)
		CollectTweetData(tweet)
	}
}

func CollectTweetData(tweet entities.Tweet) {
	// check if tweet is already in the database, update it, otherwise insert it
	if database.TweetExists(tweet.ID) {
		// Insert new public metrics
		database.InsertNewPublicMetric(&database.TweetPublicMetric{
			TweetID:      tweet.ID,
			RetweetCount: tweet.PublicMetrics.RetweetCount,
			LikeCount:    tweet.PublicMetrics.LikeCount,
			ReplyCount:   tweet.PublicMetrics.ReplyCount,
			QuoteCount:   tweet.PublicMetrics.QuoteCount,
		})
	} else {
		// Insert tweet data
		tweet_data := &database.Tweet{
			TweetedAt:     tweet.CreatedAt,
			UserTwitterID: tweet.AuthorID,
			TweetID:       tweet.ID,
			Text:          tweet.Text,
			PublicMetrics: []database.TweetPublicMetric{
				{
					TweetID:      tweet.ID,
					RetweetCount: tweet.PublicMetrics.RetweetCount,
					LikeCount:    tweet.PublicMetrics.LikeCount,
					ReplyCount:   tweet.PublicMetrics.ReplyCount,
					QuoteCount:   tweet.PublicMetrics.QuoteCount,
				},
			},
		}
		database.InsertTweet(tweet_data)
	}
}
