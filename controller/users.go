package controller

import (
	"log"
	"ServerGo/model"
)

//import "ServerGo/model"

type Users struct {

}

func (u Users) ListById(id int) []model.User {

	log.Println("Id is: %d", id)
	user := model.User{}

	users := user.ListById(id)
	/*
	resp, err := json.Marshal(user.ListById(id))
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	*/

	return users


}