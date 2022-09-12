package server_response

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"telegram-golang-bot/api"

	"telegram-golang-bot/database"
)

type ServerResponse interface {
	SendJson(tx string)
	SendText(tx string)
	SendPhoto(*os.File)
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

func (srv *ServerWT) SendPhoto(i *os.File) {
	srv.send(&Photo{
		Photo:  i,
		ChatID: srv.PrivadeMessage.Message.ChatID,
	})

}

func (srv *ServerWT) send(i Item) {
	client := &http.Client{}

	req, err := i.Request()
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

	fmt.Println(string(body))
}

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

var (
	BotKey string = "5570286790:AAH8nbOItiacjbWc7XkvhtICQCaWcmU7Aq0"
)

type Photo struct {
	Config struct {
		URL string
	}
	ChatID int
	Photo  *os.File
}

func (sp *Photo) Request() (*http.Request, error) {

	var err error

	sp.buildURL()

	payload, writer, _ := sp.bodyMime()
	req, err := http.NewRequest("GET", sp.Config.URL, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func (sp *Photo) buildURL() {
	var Paramethers []Element
	{
		chat_id := Element{Key: "chat_id", Value: sp.ChatID}
		Paramethers = []Element{chat_id}
	}

	url := ParceURL("/sendPhoto", Paramethers...)

	sp.Config.URL = url.String()
}

func (sp *Photo) bodyMime() (payload *bytes.Buffer, writer *multipart.Writer, err error) {

	// Body -->
	payload = &bytes.Buffer{}

	// MIME for Body -->
	writer = multipart.NewWriter(payload)

	defer writer.Close()

	part, err := writer.CreateFormFile("photo", filepath.Base(sp.Photo.Name()))
	if err != nil {
		return payload, writer, err
	}

	_, err = io.Copy(part, sp.Photo)
	if err != nil {
		return payload, writer, err
	}

	return payload, writer, nil
}

type Item interface {
	Request() (*http.Request, error)
	buildURL()
}

type Element struct {
	Key   string
	Value int
}

func ParceURL(method string, elem ...Element) url.URL {
	u := &url.URL{}

	u.Scheme = "https"
	u.Host = "api.telegram.org"
	u.Path = fmt.Sprintf("/bot%s%s", BotKey, method)

	q := u.Query()

	for _, v := range elem {
		q.Set(v.Key, fmt.Sprintf("%d", v.Value))
	}

	u.RawQuery = q.Encode()

	return *u

}