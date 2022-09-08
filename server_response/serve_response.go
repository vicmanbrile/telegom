package server_response

import (
	"fmt"
	"os"

	"telegram-golang-bot/api"
	"telegram-golang-bot/database"
)

type ServerResponse interface {
	SendJson(tx string)
	SendText(tx string)
}

type ServerWT struct {
	PrivadeMessage api.Update
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

//

func (srv *ServerWT) InitConversation(exists, create bool) {

	/*
		Comprovar si existe la conversación;
		(No existe) => creearla;
		(existe) => reportar;
	*/

	MC := database.NewMongoClientConversation(os.Getenv("MONGODB_CONNECTION"))

	ls, err := MC.FindConversation("6215c7dc38821f527b019d3e")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ls)

	defer MC.CancelConection()

}
