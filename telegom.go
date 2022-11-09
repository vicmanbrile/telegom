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

type HandleTelegom func(server_response.ServerWT, *api.Update)
type TeleGom struct {
	telegramToken string
	mongoToken    string
	Commands      map[string]HandleTelegom
	handlers      map[string]func(server_response.ServerResponse, *api.Update)
}

func InitTeleGom(TelegramTKN, MongoTKN string) *TeleGom {
	return &TeleGom{
		telegramToken: TelegramTKN,
		mongoToken:    MongoTKN,
		handlers: map[string]func(server_response.ServerResponse, *api.Update){
			"/help":            helpDefault,
			"/recurseNotFount": recurseNotFountDefault,
		},
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

	MGDB := database.NewMongoClientConversation(tg.mongoToken)
	defer MGDB.CancelConection()

	mssg := message.Message.Text
	// Implement detector of commands
	i := strings.Index(mssg, "/")
	if i == 0 {
		Hdr, ok := tg.handlers[mssg]

		if ok {
			WT := &server_response.ServerWT{
				PrivadeMessage: message,
			}

			Hdr(WT, &message)
		}
		// Falta un else

	} else if i <= -1 {
		ConversationContinued, err := MGDB.FindConversation(message.Message.From.ID)

		if err != nil || ConversationContinued == nil {
			fmt.Printf("Problema al buscar en base de datos: %s\n", err)
			ConversationContinued.Command = "/recurseNotFount"
		}

		// Implement Follow a Conversation
		Hdr, _ := tg.handlers[ConversationContinued.Command]

		WT := &server_response.ServerWT{
			PrivadeMessage:        message,
			ConversationContinued: *ConversationContinued,
		}

		Hdr(WT, &message)

	}

}

func (tg *TeleGom) Handle(command string, s HandleTelegom) {

	if tg.Commands == nil {
		tg.Commands = make(map[string]HandleTelegom)
	}

	switch command {
	case "/help":
		tg.handlers[command] = s
	case "/recurseNotFount":
		tg.handlers[command] = s
	default:
		tg.handlers[command] = s
	}
}

func (tg *TeleGom) ServeToTelegram(w server_response.ServerResponse, r *api.Update) {
	s, _ := tg.Commands[r.EditedMessage.Text]

	s(w, r)
}

// Handles help to user Default
func helpDefault(response server_response.ServerResponse, message *api.Update) {
	response.SendText("Use un comando")
}

func recurseNotFountDefault(response server_response.ServerResponse, message *api.Update) {
	response.SendText("No se encontro su conversacion en nuestra base de datos")
}
