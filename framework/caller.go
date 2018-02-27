package framework

import (
	"net/http"
	"fmt"
)

type handler func() string

type Caller struct {}

func (c Caller) Get(url string) {
	resp, err := http.Get(url)
	if err == nil {

		fmt.Println("Response %s", resp)
	}
}
