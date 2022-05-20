package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arshamalh/twigo"
	"github.com/arshamalh/twigo/entities"
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

var bot *twigo.Client

func main() {
	var err error

	bot, err = twigo.NewClient("", "", "", "", os.Getenv("BEARER_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	// Initialize database
	InitializeScheduler(time.UTC)

	app := fiber.New()

	app.Static("/", "./ui")

	app.Post("/tracking_users", AddTrackingUsers)

	app.Get("/tracking_users", GetTrackingUsers)

	app.Delete("/tracking_users", StopTrackingUsers)

	app.Get("/tracking_tweets", GetTrackingTweets)

	app.Delete("/tracking_tweets", StopTrackingTweets)

	log.Fatal(app.Listen(":80"))
}
