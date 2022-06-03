package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arshamalh/twigo"
	"github.com/arshamalh/twigoPanel/database"
	"github.com/gofiber/fiber/v2"
)

var bot *twigo.Client

func main() {
	var err error

	bot, err = twigo.NewClient("", "", "", "", os.Getenv("BEARER_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	database.Connect("twigo.db")

	database.AutoMigrate()

	InitializeScheduler(time.UTC)

	app := fiber.New()

	app.Static("/", "./ui")

	// Add bots, get bots, remove bot routes, instead of reading them from env.

	app.Post("/tracking_users", AddTrackingUsers)

	app.Get("/tracking_users", GetTrackingUsers)

	app.Delete("/tracking_users", StopTrackingUsers)

	app.Get("/tracking_tweets", GetTrackingTweets)

	app.Delete("/tracking_tweets", StopTrackingTweets)

	app.Get("/tweet_data", GetTweetData)

	log.Fatal(app.Listen(":80"))
}
