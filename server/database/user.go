package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserTwitterID string  `json:"user_twitter_id"`
	Username      string  `json:"username"`
	Tweets        []Tweet `json:"tweets"`
}
