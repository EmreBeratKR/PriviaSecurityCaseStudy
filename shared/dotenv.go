package shared

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func IsProductionEnvironment() bool {
	return os.Getenv("ENVIRONMENT") == "prod"
}

func IsDevelopmentEnvironment() bool {
	return !IsProductionEnvironment()
}
