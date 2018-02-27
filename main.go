package main

import (
	"net/http"
)

func error(w http.ResponseWriter, statusError int) {
	w.WriteHeader(statusError)
}

func main() {

	StartRoutes()

}