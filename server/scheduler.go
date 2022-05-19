package main

import (
	"time"

	"github.com/go-co-op/gocron"
)

var scheduler *gocron.Scheduler

func InitializeScheduler(*time.Location) {
	scheduler = gocron.NewScheduler(time.UTC)
	scheduler.StartAsync()
}
