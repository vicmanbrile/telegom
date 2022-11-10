package item_response

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)


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

