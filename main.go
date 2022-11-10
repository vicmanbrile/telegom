package main

import (
	"fmt"
	"log"
	"os"

	"telegram-golang-bot/api"

	"github.com/joho/godotenv"
)

var (
	telegramToken string
	mongoToken    string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	telegramToken = os.Getenv("TELEGRAM_TOKEN")
	mongoToken = os.Getenv("MONGODB_CONNECTION")

	var tlg = InitTeleGom()

	tlg.Handle("/start", func(response ServerResponse, update *api.Update) {

		photo, err := os.Open("./maquinados.png")
		if err != nil {
			fmt.Println(err)
		}

		defer photo.Close()

		response.SendPhoto(photo)

		response.SendText("Hola, ¿Como estas?")
	})

	Listen(tlg)

}
