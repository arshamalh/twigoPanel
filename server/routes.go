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
			msg = fmt.Sprintf("tracking user %s started", user_id)
		} else {
			msg = fmt.Sprintf("could not start tracking user %s", user_id)
		}
		json_msgs = append(json_msgs, msg)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": json_msgs,
	})
}

func StopTrackingUsers(ctx *fiber.Ctx) error {
	var user_ids []string
	var json_msgs []string

	if err := ctx.BodyParser(&user_ids); err != nil {
		return err
	}

	for _, user_id := range user_ids {
		var msg string
		if err := scheduler.RemoveByTag(fmt.Sprintf("user_track_%s", user_id)); err == nil {
			msg = fmt.Sprintf("tracking user %s stopped", user_id)
		} else {
			msg = fmt.Sprintf("could not stop tracking user %s", user_id)
		}
		json_msgs = append(json_msgs, msg)
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": json_msgs,
	})
}

func StopTrackingTweets(ctx *fiber.Ctx) error {
	var tweet_ids []string
	var json_msgs []string

	if err := ctx.BodyParser(&tweet_ids); err != nil {
		return err
	}

	for _, tweet_id := range tweet_ids {
		var msg string
		if err := scheduler.RemoveByTag(fmt.Sprintf("tweet_track_%s", tweet_id)); err == nil {
			msg = fmt.Sprintf("tracking tweet %s stopped", tweet_id)
		} else {
			msg = fmt.Sprintf("could not stop tracking tweet %s", tweet_id)
		}
		json_msgs = append(json_msgs, msg)
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": json_msgs,
	})
}
