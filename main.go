package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tlg := InitTeleGom(os.Getenv("TELEGRAM_TOKEN"))

	tlg.CancelForCommand("/start")
}
