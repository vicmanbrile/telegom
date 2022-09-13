package server_response

import (
	"fmt"
	"os"
	"telegram-golang-bot/api"
	items "telegram-golang-bot/server_response/items"

	"telegram-golang-bot/database"
)

type ServerResponse interface {
	SendJson(tx string)
	SendText(tx string)
	SendPhoto(*os.File)
}

type ServerWT struct {
	PrivadeMessage  api.Update
	CommandPendient database.CommandsPending
}

/*

func (srv *ServerWT) newPending(command string, cp CommandsPending) {
	srv.Pending[command] = cp
}

func (srv *ServerWT) cancelPending(key string) {
	_, ok := srv.Pending[key]
	if ok {
		delete(srv.Pending, key)
	}
}

*/

func (srv *ServerWT) SendJson(tx string) {
	fmt.Println(tx)
}

func (srv *ServerWT) SendText(tx string) {
	fmt.Println(tx)
}

func (srv *ServerWT) SendPhoto(i *os.File) {
	items.Send(&items.Photo{
		Photo:  i,
		ChatID: srv.PrivadeMessage.Message.From.ID,
	})

}

func (srv *ServerWT) InitConversation(exists, create bool) {

	/*
		Comprovar si existe la conversación;
		(No existe) => creearla;
		(existe) => reportar;
	*/

	MC := database.NewMongoClientConversation(os.Getenv("MONGODB_CONNECTION"))

	ls, err := MC.FindConversation(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ls)

	defer MC.CancelConection()

}
