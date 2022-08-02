package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (tg *TeleGom) NewPendient(command string, cp *CommandsPending) {
	tg.Pendients[command] = *cp
}

func (tg *TeleGom) CancelForCommand(key string) {
	_, ok := tg.Pendients[key]
	if ok {
		delete(tg.Pendients, key)
	}
}

func (tg *TeleGom) Get() {
	AvailableMethod := "getMe"

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", tg.TelegramToken, AvailableMethod)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var jsonResp api.Bot

	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
