package item_response

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
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
	u.Path = fmt.Sprintf("/bot%s%s", os.Getenv("TELEGRAM_TOKEN"), method)

	q := u.Query()

	for _, v := range elem {
		q.Set(v.Key, fmt.Sprintf("%d", v.Value))
	}

	u.RawQuery = q.Encode()

	return *u

}

func Send(i Item) {
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

