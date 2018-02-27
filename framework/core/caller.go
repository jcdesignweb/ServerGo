package core

import (
	"net/http"
	"io/ioutil"
)

type handler func(response string) string

type Build interface {
	Make()
}

type Caller struct {}

func (c Caller) Get(url string, callback handler) string{
	response, err := http.Get(url)
	if err == nil {
		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err == nil {
			bodyString := string(bodyBytes)
			return callback(bodyString)
		}
	}

	return ""
}
