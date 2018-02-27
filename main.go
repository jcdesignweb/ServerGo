package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"ServerGo/controller"
	"strconv"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		if id, ok := vars["id"]; ok {

			idInt, err := strconv.Atoi(id)
			if err == nil {
				var usersController = controller.Users{}
				usersController.ListById(idInt)

				//fmt.Fprintf(w, "ID: %v\n", vars["id"])
				return
			}
		}

		error(w, http.StatusBadRequest)

	}).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:9000", router))
}

func error(w http.ResponseWriter, statusError int) {
	w.WriteHeader(statusError)
}

func main() {

	handleRequests()
}