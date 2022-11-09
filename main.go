package main

import (
	"fmt"
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

	var tlg = InitTeleGom(os.Getenv("TELEGRAM_TOKEN"), os.Getenv("MONGODB_CONNECTION"))

	tlg.Handle("/start", func(response server_response.ServerResponse, update *api.Update) {

		photo, err := os.Open("./maquinados.png")
		if err != nil {
			fmt.Println(err)
		}

		defer photo.Close()

		response.SendPhoto(photo)

		response.SendText("Hola, ¿Como estas?")
	})

	Listen(tlg)

	tlg.Listen()
}

type ServerTelegom interface {
	ServeToTelegram(server_response.ServerResponse, *api.Update)
}

// Se busca la manera de resivir y responder

func Listen(s ServerTelegom) {

	Maria := &server_response.ServerResponse{}

	/*
		ideas:
		- Entregar ChatID para poder buscar en base de datos
		- Todos los metodos deben de retornar error
		- El packete no imprime nada y deja que el desarollador imprima en su consola o al usuario
	*/

	// Se cargan los metodos de envio
	SS := &server_response.ServerWT{}

	s.ServeToTelegram(SS, Maria)

}
