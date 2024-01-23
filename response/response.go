package response

import (
	"fmt"
	"os"
	items "telegom/response/items"

	"telegom/database"
)

type Response struct {
	FromID int
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
		FromID: srv.FromID,
	})

}

func (srv *Response) InitConversation(exists, create bool) {

	/*
		Comprovar si existe la conversación;
		(No existe) => creearla;
		(existe) => reportar;
	*/

	MC := database.NewMongoClientConversation(os.Getenv("MONGODB_CONNECTION"))

	ls, err := MC.FindConversation(srv.FromID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ls)

	defer MC.CancelConection()

}
