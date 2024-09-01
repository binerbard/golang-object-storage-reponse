package utilssetting

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// CastInt : Function for casting string to int
func CastInt(val string) int {
	result, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Error casting %s to int: %s", val, err)
	}
	return result
}

// GetEnv : Function for get env
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// LoadEnv : Function for load env
func LoadEnv() {
	if os.Getenv("GO_ENV") != "production" {
	if err := godotenv.Load(".env"); err != nil {
		if err := godotenv.Load("../../.env"); err != nil {
			log.Printf("Warning: Error loading .env file: %s", err)
		}
	}
	}
}