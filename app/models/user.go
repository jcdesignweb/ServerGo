package models

import (
	"fmt"
	"ServerGo/app/models/request"
	"ServerGo/boot"
)

// Contact struct
type User struct {
	id          int          `json:"id"`
	fistName        string       `json:"first_name"`
	lastName        string       `json:"last_name"`
	avatar       string       `json:"avatar"`

}

// Contacts slice of contact
type Users []*User

var application = boot.App

func (u *User) GetUsers(id int) string {

	var method = "api/users"

	var url = fmt.Sprintf("%s%s?page=%d", application.UriApi, method, id)

	req := request.Request{}
	usersJSON, status := req.Get(url)

	if status {
		// Maybe here I need make parse JSON with My User struct. But in this case this is not necessary
		return usersJSON
	}

	return ""

}

// String return the string of object
func (u *User) String() string {
	return fmt.Sprintf("%d, %s, %s, %s", u.id, u.fistName, u.lastName, u.avatar)
}

// Parse json response and return an User, for now this return empty User object
func (u *User) parse(response string) (user User) {
	return User{}
}
