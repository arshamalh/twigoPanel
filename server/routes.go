package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AddTrackingUsers(ctx *fiber.Ctx) error {
	var user_ids []string
	var json_msgs []string

	if err := ctx.BodyParser(&user_ids); err != nil {
		return err
	}

	for _, user_id := range user_ids {
		job_name := fmt.Sprintf("user_track_%s", user_id)
		_, err := scheduler.Every(5).Minutes().Tag(job_name).Do(TrackUser, user_id)
		var msg string
		if err == nil {
			msg = fmt.Sprintf("tracking user %s started\n", user_id)
		} else {
			msg = fmt.Sprintf("could not start tracking user %s", user_id)
		}
		json_msgs = append(json_msgs, msg)
	}

	return ctx.Status(200).JSON(json_msgs)
}

func StopTrackingUsers(ctx *fiber.Ctx) error {
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
}

func StopTrackingTweets(ctx *fiber.Ctx) error {
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
}
