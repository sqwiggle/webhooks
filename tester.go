package main

import(
	"net/http"
	"strings"
	"io/ioutil"
)

type Tester struct {
	Url string
	Data string
}

func (t *Tester) Execute () (*Attempt, error) {
	resp, err := http.Post(t.Url, "test/html", strings.NewReader(t.Data))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Attempt{
		Response:string(body),
		Status:resp.StatusCode,
	}, nil
}

