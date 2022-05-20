package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

func Connect(dsn string) {
	var err error
	DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
}

func AutoMigrate() {
	err := DBConnection.AutoMigrate(User{}, Tweet{})
	if err != nil {
		return
	}
	fmt.Println("Database Migrated")
}
