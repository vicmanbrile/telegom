package main

import (
	"fmt"
	"strings"
	"telegram-golang-bot/database"
	"telegram-golang-bot/response"

	"telegram-golang-bot/api"
)

type HandleTelegom func(ServerResponse, *api.Update)

type TeleGom struct {
	handlers map[string]HandleTelegom
}

func InitTeleGom() *TeleGom {
	return &TeleGom{
		handlers: map[string]HandleTelegom{
			"/help":            helpDefault,
			"/recurseNotFount": recurseNotFountDefault,
		},
	}
}

func (tg *TeleGom) Handle(command string, s HandleTelegom) {

	if tg.handlers == nil {
		tg.handlers = make(map[string]HandleTelegom)
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

func (tg *TeleGom) ServeToTelegram(w ServerResponse, r *api.Update) {

	MGDB := database.NewMongoClientConversation(mongoToken)
	defer MGDB.CancelConection()

	mssg := r.Message.Text
	// Implement detector of commands
	indexOfChater := strings.Index(mssg, "/")
	if indexOfChater == 0 {
		handler, ok := tg.handlers[mssg]

		if ok {
			res := &response.Response{
				ChatID: r.Message.ChatID,
			}

			handler(res, r)
		} else {
			res := &response.Response{
				ChatID: r.Message.ChatID,
			}

			tg.handlers["/help"](res, r)
		}

	} else if indexOfChater <= -1 {
		ConversationContinued, err := MGDB.FindConversation(r.Message.From.ID)

		if err != nil || ConversationContinued == nil {
			fmt.Printf("Problema al buscar en base de datos: %s\n", err)
			ConversationContinued.Command = "/recurseNotFount"
		}

		// Implement Follow a Conversation
		Hdr, _ := tg.handlers[ConversationContinued.Command]

		WT := &response.Response{
			ChatID: *&r.Message.ChatID,
		}

		Hdr(WT, r)

	} else {
		// Si el "/" se encuentra en el mensaje eg.{"Na/Sodio", "Amigo/Friend"}
		s, _ := tg.handlers[r.EditedMessage.Text]

		s(w, r)
	}

}
