package common

import "os"

func IsProductionEnvironment() bool {
	return os.Getenv("ENVIRONMENT") == "prod"
}

func IsDevelopmentEnvironment() bool {
	return !IsProductionEnvironment()
}
