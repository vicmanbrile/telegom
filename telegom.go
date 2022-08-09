package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitlab.com/vicmanbrile/telegram-golang-bot/api"
)

// CommandsPending
// Can be canceled by two events
// CancelForCommand(), CancelForTime()

type CommandsPending struct {
	Steps   int
	Process int
	Data    map[string]string
}

type TeleGom struct {
	telegramToken string
	pendients     map[string]CommandsPending
	handlers      map[string]func(ServerResponse, *api.Update)
}

func InitTeleGom(TelegramTKN string) *TeleGom {
	return &TeleGom{
		telegramToken: TelegramTKN,
		pendients:     map[string]CommandsPending{},
		handlers:      map[string]func(ServerResponse, *api.Update){},
	}
}

func (tg *TeleGom) Listen() {

	var offSet int
	var status bool

	status = true

	for status {
		// Robot method ("getMe")
		result := tg.get("getUpdates", fmt.Sprintf("?offset=%d", offSet))

		for _, v := range result.Update {
			offSet = v.UpdateID + 1
			fmt.Printf("Offset: %d\n", offSet)
			tg.SendMessage(v)
		}

	}

}

func (tg *TeleGom) newPendient(command string, cp *CommandsPending) {
	tg.pendients[command] = *cp
}

func (tg *TeleGom) cancelPendient(key string) {
	_, ok := tg.pendients[key]
	if ok {
		delete(tg.pendients, key)
	}
}

func (tg *TeleGom) get(AvailableMethod string, parameter string) *api.Updates {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s%s", tg.telegramToken, AvailableMethod, parameter)

	// Http Cliend to String
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

type ServerResponse interface {
	SendJson(tx string)
}

type ServerRS struct{}

func (m *ServerRS) SendJson(tx string) {
	fmt.Println(tx)
}

func (tg *TeleGom) Handle(command string, s func(ServerResponse, *api.Update)) {
	tg.handlers[command] = s
}

func (tg *TeleGom) SendMessage(update api.Update) {

	// Implement detector of commands
	Hdr, ok := tg.handlers[update.Message.Text]
	if ok {
		if _, ok := tg.pendients[update.Message.ChatID]; ok {
		}
	}
	// Implement Folow a Conversation

	MT := &ServerRS{}

	Hdr(MT, &update)
}
