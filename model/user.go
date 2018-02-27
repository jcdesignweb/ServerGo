package model

import (
	"fmt"
	"ServerGo/config"
	"log"
	"ServerGo/framework/core"
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
	log.Println(url)
	caller := core.Caller{}


	users := caller.Get(url, func(response string) string {
		return response
	})

	fmt.Println(users)

	return []User {
		User{}, User{},
	}

}

