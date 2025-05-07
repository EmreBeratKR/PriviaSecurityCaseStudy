package shared

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file... Skipping...")
	}
}

func IsProductionEnvironment() bool {
	return os.Getenv("ENVIRONMENT") == "prod"
}

func IsDevelopmentEnvironment() bool {
	return !IsProductionEnvironment()
}
