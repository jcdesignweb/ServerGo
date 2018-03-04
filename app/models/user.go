package models

import (
	"fmt"
	"ServerGo/boot"
	"encoding/json"
	"net/http"
)

// User struct
type User struct {
	ID          int    	`json:"id"`
	FistName	string	`json:"first_name"`
	LastName    string	`json:"last_name"`
	Avatar      string	`json:"avatar"`
}

type UserResponse struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []User `json:"data"`
}

type Users []*User

var application = boot.App

func (u *User) GetUsers(id int) (UserResponse, error) {
	var response UserResponse
	var method = "api/users"
	var url = fmt.Sprintf("%s%s?page=%d", application.UriApi, method, id)

	request, err := http.Get(url)
	if err == nil {

		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&response)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("whoops:", err)
		}
	}

	return response, err
}

// String return the string of object
func (u *User) String() string {
	return fmt.Sprintf("%d, %s, %s, %s", u.ID, u.FistName, u.LastName, u.Avatar)
}

