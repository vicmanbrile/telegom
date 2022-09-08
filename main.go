package main

import (
	"log"
	"os"

	"telegram-golang-bot/api"
	"telegram-golang-bot/server_response"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var tlg = InitTeleGom(os.Getenv("TELEGRAM_TOKEN"))

	tlg.Handle("/start", func(response server_response.ServerResponse, update *api.Update) {

		response.SendText("Hola, ¿Como estas?")
	})

	tlg.Listen()
}
