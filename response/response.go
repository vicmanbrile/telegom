package response

import (
	"fmt"
	"os"
	items "telegram-golang-bot/response/items"

	"telegram-golang-bot/database"
)

type Response struct {
	ChatID int
}

func (srv *Response) SendJson(tx string) {
	fmt.Println(tx)
}

func (srv *Response) SendText(tx string) {
	fmt.Println(tx)
}

func (srv *Response) SendPhoto(i *os.File) {
	items.Send(&items.Photo{
		Photo:  i,
		ChatID: srv.ChatID,
	})

}

func (srv *Response) InitConversation(exists, create bool) {

	/*
		Comprovar si existe la conversación;
		(No existe) => creearla;
		(existe) => reportar;
	*/

	MC := database.NewMongoClientConversation(os.Getenv("MONGODB_CONNECTION"))

	ls, err := MC.FindConversation(srv.ChatID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ls)

	defer MC.CancelConection()

}
