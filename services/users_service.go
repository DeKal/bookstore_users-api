package services

import (
	userdao "github.com/DeKal/bookstore_users-api/domain/users/dao"
	userdto "github.com/DeKal/bookstore_users-api/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/utils/dates"
	"github.com/DeKal/bookstore_users-api/utils/errors"
)

// CreateUser create a user in DB
func CreateUser(user userdto.User) (*userdto.User, *errors.RestError) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	user.Status = userdto.StatusActive
	user.DateCreated = dates.GetNowDBString()

	err = userdao.Save(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser get a user in DB
func GetUser(userID int64) (*userdto.User, *errors.RestError) {
	if userID <= 0 {
		return nil, errors.NewBadRequestError("Invalid user id")
	}

	target := &userdto.User{
		ID: userID,
	}
	err := userdao.Get(target)
	if err != nil {
		return nil, err
	}

	return target, nil
}

// UpdateUser update a user in DB
func UpdateUser(user userdto.User) (*userdto.User, *errors.RestError) {
	existedUser, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	existedUser.FirstName = user.FirstName
	existedUser.LastName = user.LastName
	existedUser.Email = user.Email

	err = userdao.Update(existedUser)
	if err != nil {
		return nil, err
	}

	return existedUser, nil
}

// PatchUser update a user in DB
func PatchUser(user userdto.User) (*userdto.User, *errors.RestError) {
	existedUser, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if user.FirstName != "" {
		existedUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		existedUser.LastName = user.LastName
	}
	if user.Email != "" {
		existedUser.Email = user.Email
	}

	err = userdao.Update(existedUser)
	if err != nil {
		return nil, err
	}

	return existedUser, nil
}

// DeleteUser to Delete a user with given userId
func DeleteUser(userID int64) *errors.RestError {
	if userID <= 0 {
		return errors.NewBadRequestError("Invalid user id")
	}

	err := userdao.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}

// Search to search users having given status
func Search(status string) ([]userdto.User, *errors.RestError) {
	if status == "" {
		return nil, errors.NewBadRequestError("Invalid status")
	}
	users, err := userdao.FindByStatus(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error)
	}

	return users, nil
}
