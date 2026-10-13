package application

import (
	"fmt"
	"os"
	"strconv"
)

type config struct {
	databaseURL string

	redisAddress string // optional redis address
	serverPort   uint16 // optional server port
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

	cfg := config{
		databaseURL: dbURL,

		redisAddress: "localhost:6379",
		serverPort:   8080,
	}

	if redisAddr, ok := os.LookupEnv("REDIS_ADDR"); ok {
		cfg.redisAddress = redisAddr
	}

	if serverPort, ok := os.LookupEnv("SERVER_PORT"); ok {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.serverPort = uint16(port)
		}
	}

	return cfg
}
