package request

import (
	"net/http"
	"io/ioutil"
)

type Request struct {}

/**
 * Make an GET request
 */
func (r Request) Get(url string) (string, bool) {
	response, err := http.Get(url)
	if err == nil {
		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err == nil {
			bodyString := string(bodyBytes)
			return bodyString, true
		}
	}

	return "Error", false
}
