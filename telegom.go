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
}

type TeleGom struct {
	TelegramToken string
	Pendients     map[string]CommandsPending
}

func InitTeleGom(TelegramTKN string) *TeleGom {
	return &TeleGom{
		TelegramToken: TelegramTKN,
		Pendients:     map[string]CommandsPending{},
	}
}

func (tg *TeleGom) Listen() {

	var offSet int
	var status bool

	status = true

	for status {
		// Robot method ("getMe")
		result := tg.Get("getUpdates", fmt.Sprintf("?offset=%d", offSet))

		for _, v := range result.Update {
			offSet = v.UpdateID + 1
			fmt.Printf("Offset: %d\n", offSet)
			SendMessage(v)
		}

	}

}

func (tg *TeleGom) NewPendient(command string, cp *CommandsPending) {
	tg.Pendients[command] = *cp
}

func (tg *TeleGom) CancelPendient(key string) {
	_, ok := tg.Pendients[key]
	if ok {
		delete(tg.Pendients, key)
	}
}

func (tg *TeleGom) Get(AvailableMethod string, parameter string) *api.Updates {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s%s", tg.TelegramToken, AvailableMethod, parameter)

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

func SendMessage(update api.Update) {
	fmt.Println(update.Message.ReplyToMessageID)
}