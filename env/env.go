package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load loads environment variables from .env file if present.
func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[shared-go] .env file not found or skipped (using system env)")
	}
}

// Get returns an environment variable or a fallback if empty.
func Get(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
