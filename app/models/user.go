package models

import "fmt"

// Contact struct
type User struct {
	id          int          `json:"id"`
	fistName        string       `json:"first_name"`
	lastName        string       `json:"last_name"`
	avatar       string       `json:"avatar"`

}

// Contacts slice of contact
type Users []*User

// String return the string of object
func (c *User) String() string {
	return fmt.Sprintf("%d, %s, %s, %s", c.id, c.fistName, c.lastName, c.avatar)
}

// String of contacts
func (cs *Users) String() string {
	var r string
	for _, c := range *cs {
		r += fmt.Sprintf("%d, %s, %s, %s\n", c.id, c.fistName, c.lastName, c.avatar)
	}
	return r
}