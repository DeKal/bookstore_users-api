package services

import (
	"github.com/DeKal/bookstore_users-api/domain/users"
	"github.com/DeKal/bookstore_users-api/utils/errors"
)

// CreateUser create a user in DB
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	err = user.Save()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser get a user in DB
func GetUser(userID int64) (*users.User, *errors.RestError) {
	if userID <= 0 {
		return nil, errors.NewBadRequestError("Invalid user id")
	}

	target := users.User{
		ID: userID,
	}
	err := target.Get()
	if err != nil {
		return nil, err
	}

	return &target, nil
}
