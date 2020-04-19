package userdto

import (
	"strings"

	"github.com/DeKal/bookstore_users-api/src/utils/errors"
)

const (
	// StatusActive status of user is active
	StatusActive = "active"
	// StatusInActive status of user is inactive
	StatusInActive = "inactive"
)

// User contains information of Users
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Users is slice of User
type Users []User

// Validate check if user info is valid
func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address ")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}
