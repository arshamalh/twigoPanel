package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
}

func AutoMigrate() {
	err := DB.AutoMigrate(User{}, Tweet{}, TweetPublicMetrics{})
	if err != nil {
		return
	}
	fmt.Println("Database Migrated")
}
