package main

import (
	"net/http"
	"ServerGo/config"
	"github.com/gorilla/mux"
	"strconv"
	"ServerGo/controller"
	"log"
	"fmt"
)

/**
 * Routing
 */
func StartRoutes() {
	config := config.Configuration{}

	r := mux.NewRouter()
	r.Methods("GET", "POST", "DELETE")
	r.Headers("X-Requested-With", "XMLHttpRequest")

	router := r.StrictSlash(true)
	sub := router.PathPrefix(config.GetPrefixApi()).Subrouter()
	sub.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		if id, ok := vars["id"]; ok {

			idInt, err := strconv.Atoi(id)
			if err == nil {
				var usersController = controller.Users{}
				usersController.ListById(idInt)
				return
			}
		}

		error(w, http.StatusBadRequest)

	}).Methods("GET")

	host := fmt.Sprintf("%s:%s", config.GetHost(), config.GetPort())

	log.Fatal(http.ListenAndServe(host, router))
}
