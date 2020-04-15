package users

import (
	"fmt"

	"github.com/DeKal/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

// Save to persist User to DB
func (user *User) Save() *errors.RestError {
	findUser := userDB[user.ID]
	if findUser != nil {
		errorMsg := fmt.Sprintf("User %d has already registered", user.ID)
		if findUser.Email == user.Email {
			errorMsg = fmt.Sprintf("User email %s has already existed", user.Email)
		}
		return errors.NewBadRequestError(errorMsg)
	}
	userDB[user.ID] = user
	return nil
}

// Get to get User from DB
func (user *User) Get() *errors.RestError {
	target := userDB[user.ID]
	if target == nil {
		errorMsg := fmt.Sprintf("User %d not found", user.ID)
		return errors.NewNotFoundError(errorMsg)
	}

	user.ID = target.ID
	user.Email = target.Email
	user.FirstName = target.FirstName
	user.LastName = target.LastName
	user.DateCreated = target.DateCreated

	return nil
}
