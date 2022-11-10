package main

import "telegram-golang-bot/api"

// Handles help to user Default
func helpDefault(response ServerResponse, message *api.Update) {
	response.SendText("Use Command")
}

func recurseNotFountDefault(response ServerResponse, message *api.Update) {
	response.SendText("Your conversation was not found in our database")
}
