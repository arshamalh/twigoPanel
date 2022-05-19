package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/arshamalh/twigo"
	"github.com/arshamalh/twigo/entities"
	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string
}

type TweetData struct {
	CreatedAt     time.Time                             `json:"created_at"`
	AuthorID      string                                `json:"author_id"`
	ID            string                                `json:"id"`
	Text          string                                `json:"text"`
	PublicMetrics map[int64]entities.TweetPublicMetrics `json:"public_metrics"`
}

var tweet_authors = map[string]string{}
var found_tweets = []string{}
var bot *twigo.Client
var scheduler *gocron.Scheduler

func main() {
	var err error

	bearer_token := flag.String("bearer", "", "Bearer token")
	flag.Parse()

	bot, err = twigo.NewClient("", "", "", "", *bearer_token)
	if err != nil {
		fmt.Println(err)
	}

	// Initialize database

	scheduler = gocron.NewScheduler(time.UTC)

	app := fiber.New()

	app.Post("/add_users_to_be_tracked", func(ctx *fiber.Ctx) error {
		var user_ids []string

		if err := ctx.BodyParser(&user_ids); err != nil {
			return err
		}

		for _, user_id := range user_ids {
			_, err := scheduler.Every(5).Minutes().Tag(fmt.Sprintf("user_track_%s", user_id)).Do(TrackUser, user_id)
			if err == nil {
				fmt.Printf("tracking user %s started\n", user_id)
			} else {
				fmt.Println(err)
			}
		}

		return ctx.JSON(true)
	})

	app.Post("/stop_tracking_users", func(ctx *fiber.Ctx) error {
		var user_ids []string

		if err := ctx.BodyParser(&user_ids); err != nil {
			return err
		}

		for _, user_id := range user_ids {
			if err := scheduler.RemoveByTag(fmt.Sprintf("user_track_%s", user_id)); err == nil {
				fmt.Printf("tracking user %s stopped\n", user_id)
			}
		}
		return ctx.JSON(true)
	})

	app.Post("/stop_tracking_tweets", func(ctx *fiber.Ctx) error {
		var tweet_ids []string

		if err := ctx.BodyParser(&tweet_ids); err != nil {
			return err
		}

		for _, tweet_id := range tweet_ids {
			if err := scheduler.RemoveByTag(fmt.Sprintf("tweet_track_%s", tweet_id)); err == nil {
				fmt.Printf("tracking tweet %s stopped\n", tweet_id)
			}
		}
		return ctx.JSON(true)
	})

	scheduler.StartAsync()
	log.Fatal(app.Listen(":3000"))
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

func contains(arrayOfStrings []string, string_item string) bool {
	for _, val := range arrayOfStrings {
		if val == string_item {
			return true
		}
	}
	return false
}
