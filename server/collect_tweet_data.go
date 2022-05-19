package main

import (
	"fmt"
	"time"

	"github.com/arshamalh/twigo/entities"
)

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
