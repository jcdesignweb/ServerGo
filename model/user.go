package model

import (
	"net/http"
	"fmt"
	"ServerGo/config"
	"ServerGo/framework"
)

type User struct {
	id int
	firstName string
	lastName string
	avatar string
}

func (u User) ListById(id int) []User {
	config := config.Configuration{}
	var url = fmt.Sprintf("%s?page=%d", config.GetBaseUrl(), id)

	caller := framework.Caller{}
	caller.Get(url)

	var users = []User {
		User{}, User{},
	}

	return users
}

func (u User) make() {

}

