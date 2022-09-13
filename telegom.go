package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"telegram-golang-bot/database"
	"telegram-golang-bot/server_response"

	"telegram-golang-bot/api"
)

type TeleGom struct {
	telegramToken string
	mongoToken    string
	handlers      map[string]func(server_response.ServerResponse, *api.Update)
}

func InitTeleGom(TelegramTKN, MongoTKN string) *TeleGom {
	return &TeleGom{
		telegramToken: TelegramTKN,
		mongoToken:    MongoTKN,
		handlers:      map[string]func(server_response.ServerResponse, *api.Update){},
	}
}

func (tg *TeleGom) Listen() {

	var offSet int

	status := true

	for status {
		result := tg.getUpdates(fmt.Sprintf("?offset=%d", offSet))

		for _, v := range result.Update {
			offSet = v.UpdateID + 1
			tg.responseMessage(v)
		}

	}

}

func (tg *TeleGom) getUpdates(parameter string) *api.Updates {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates%s", tg.telegramToken, parameter)

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

func (tg *TeleGom) responseMessage(message api.Update) {

	// Implement detector of commands
	chatId := message.Message.From.ID

	MGDB := database.NewMongoClientConversation(tg.mongoToken)
	// defer MGDB.CancelConection()

	CommandPendient, _ := MGDB.FindConversation(chatId)

	mssg := message.Message.Text

	i := strings.Index(mssg, "/")
	if i == 0 {
		Hdr, ok := tg.handlers[mssg]

		if ok {
			// Implement Follow a Conversation
			WT := &server_response.ServerWT{
				PrivadeMessage:  message,
				CommandPendient: *CommandPendient,
			}

			Hdr(WT, &message)
		}

	} else if i <= -1 {
		Hdr, ok := tg.handlers[CommandPendient.Command]

		if ok {
			// Implement Follow a Conversation
			WT := &server_response.ServerWT{
				PrivadeMessage:  message,
				CommandPendient: *CommandPendient,
			}

			Hdr(WT, &message)
		}
	}

}

func (tg *TeleGom) Handle(command string, s func(server_response.ServerResponse, *api.Update)) {
	tg.handlers[command] = s
}
