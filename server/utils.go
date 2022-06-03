package main

import "fmt"

func remove(slice []string, s string) []string {
	new_slice := []string{}
	for _, v := range slice {
		if v != s {
			new_slice = append(new_slice, v)
		}
	}
	return new_slice
}

func contains(arrayOfStrings []string, string_item string) bool {
	for _, val := range arrayOfStrings {
		if val == string_item {
			return true
		}
	}
	return false
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  bool
	TimeZone string
}

func (dbc DBConfig) String() string {
	if dbc.Host == "" {
		dbc.Host = "localhost"
	}

	if dbc.Port == 0 {
		dbc.Port = 5432
	}

	var sslmode string
	if dbc.SSLMode {
		sslmode = "enable"
	} else {
		sslmode = "disable"
	}

	if dbc.TimeZone == "" {
		dbc.TimeZone = "Asia/Tehran"
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbc.Host,
		dbc.User,
		dbc.Password,
		dbc.DBName,
		dbc.Port,
		sslmode,
		dbc.TimeZone,
	)
}
