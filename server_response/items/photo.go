package item_response

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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
	var Parameters []Element
	{
		chatId := Element{Key: "chat_id", Value: sp.ChatID}
		Parameters = []Element{chatId}
	}

	newUrl := ParceURL("/sendPhoto", Parameters...)

	sp.Config.URL = newUrl.String()
}

func (sp *Photo) bodyMime() (payload *bytes.Buffer, writer *multipart.Writer, err error) {

	// Body -->
	payload = &bytes.Buffer{}

	// MIME for Body -->
	writer = multipart.NewWriter(payload)

	defer func(writer *multipart.Writer) {
		err := writer.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(writer)

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
