package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"telegram-golang-bot/api"
	"telegram-golang-bot/response"
)

type ServerResponse interface {
	SendJson(tx string)
	SendText(tx string)
	SendPhoto(*os.File)
}

type ServerTelegom interface {
	ServeToTelegram(ServerResponse, *api.Update)
}

func Listen(s ServerTelegom) {

	var offSet int

	status := true

	for status {
		result := getUpdates(fmt.Sprintf("?offset=%d", offSet))

		for _, update := range result.Update {
			offSet = update.UpdateID + 1

			// Se cargan los metodos de envio

			clientMessage := &response.Response{
				FromID: update.Message.From.ID,
			}

			s.ServeToTelegram(clientMessage, &update)
		}

	}

}

/*
	getUpdates() Obtiene todos los mensajes enviados al bot
*/

func getUpdates(parameter string) *api.Updates {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates%s", telegramToken, parameter)

	// Http Client to String
	Client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != err {
		fmt.Println(err)
	}

	response, err := Client.Do(request)
	if err != err {
		fmt.Println(err)
	}

	defer response.Body.Close()

	Result, err := io.ReadAll(response.Body)
	if err != err {
		fmt.Println(err)
	}
	// Close Http Client
	var jsonResp api.Updates
	err = json.Unmarshal(Result, &jsonResp)

	if err != err {
		fmt.Println(err)
	}
	return &jsonResp
}
