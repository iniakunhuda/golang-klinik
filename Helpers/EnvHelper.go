package Helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
