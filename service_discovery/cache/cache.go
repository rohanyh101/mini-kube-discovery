package cache

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Init() {

	if os.Getenv("ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		log.Println("Using development env variables")
	} else {
		log.Println("Using production env variables")
	}

	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB: func() int {
			dbNumber, err := strconv.Atoi(os.Getenv("REDIS_DB"))
			if err != nil {
				log.Fatalf("Failed to convert REDIS_DB to int: %v", err)
			}

			return dbNumber
		}(),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// check connection via pinging the server
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	log.Println("Connected to redis server successfully")
}
