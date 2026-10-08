package application

import (
	"fmt"
	"os"
)

type config struct {
	databaseURL string
}

func LoadConfig() config {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		dbUser,
		dbPass,
		dbHost,
		dbName,
	)

	return config{
		databaseURL: dbURL,
	}
}
