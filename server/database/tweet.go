package database

import (
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model                        // Created at is referening to "record" creation date!
	TweetID       string              `json:"tweet_id"`
	TweetedAt     time.Time           `json:"tweeted_at"`
	UserTwitterID string              `json:"author_id"`
	UserID        uint                `json:"user_id"`
	Text          string              `json:"text"`
	PublicMetrics []TweetPublicMetric `json:"public_metrics"`
}

type TweetPublicMetric struct {
	gorm.Model
	TweetID      string `json:"tweet_id"`
	RetweetCount int    `json:"retweet_count"`
	ReplyCount   int    `json:"reply_count"`
	LikeCount    int    `json:"like_count"`
	QuoteCount   int    `json:"quote_count"`
}

func InsertTweet(tweet *Tweet) {
	DB.Create(&tweet)
}

func InsertNewPublicMetric(metric *TweetPublicMetric) {
	DB.Create(&metric)
}

func TweetExists(tweet_id string) bool {
	var tweet Tweet
	DB.Where("tweet_id = ?", tweet_id).First(&tweet)
	return tweet.ID != 0
}

func GetTweet(id string) Tweet {
	var tweet Tweet
	DB.Preload("TweetPublicMetric").First(&tweet, id)
	return tweet
}

func GetTweets(ids []string) []Tweet {
	var tweets []Tweet
	DB.Preload("TweetPublicMetric").Where("tweet_id IN (?)", ids).Find(&tweets)
	return tweets
}
