package services

import (
	userdao "github.com/DeKal/bookstore_users-api/src/domain/users/dao"
	userdto "github.com/DeKal/bookstore_users-api/src/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/src/utils/crypto"
	"github.com/DeKal/bookstore_users-api/src/utils/dates"
	"github.com/DeKal/bookstore_users-api/src/utils/errors"
)

var (
	// UsersService is service for Users
	UsersService usersServiceInterface = &usersService{}
	userDAO                            = userdao.UserDAO
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(userdto.User) (*userdto.User, *errors.RestError)
	GetUser(int64) (*userdto.User, *errors.RestError)
	UpdateUser(userdto.User) (*userdto.User, *errors.RestError)
	PatchUser(userdto.User) (*userdto.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	Search(string) (userdto.Users, *errors.RestError)
	Login(userdto.LoginRequest) (*userdto.User, *errors.RestError)
}

// CreateUser create a user in DB
func (*usersService) CreateUser(user userdto.User) (*userdto.User, *errors.RestError) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	user.Status = userdto.StatusActive
	user.DateCreated = dates.GetNowDBString()
	user.Password = crypto.GetMD5(user.Password)

	err = userDAO.Save(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser get a user in DB
func (*usersService) GetUser(userID int64) (*userdto.User, *errors.RestError) {
	if userID <= 0 {
		return nil, errors.NewBadRequestError("Invalid user id")
	}

	target := &userdto.User{
		ID: userID,
	}
	err := userDAO.Get(target)
	if err != nil {
		return nil, err
	}

	return target, nil
}

// UpdateUser update a user in DB
func (*usersService) UpdateUser(user userdto.User) (*userdto.User, *errors.RestError) {
	existedUser, err := UsersService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	existedUser.FirstName = user.FirstName
	existedUser.LastName = user.LastName
	existedUser.Email = user.Email

	err = userDAO.Update(existedUser)
	if err != nil {
		return nil, err
	}

	return existedUser, nil
}

// PatchUser update a user in DB
func (*usersService) PatchUser(user userdto.User) (*userdto.User, *errors.RestError) {
	existedUser, err := UsersService.GetUser(user.ID)
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

	err = userDAO.Update(existedUser)
	if err != nil {
		return nil, err
	}

	return existedUser, nil
}

// DeleteUser to Delete a user with given userId
func (*usersService) DeleteUser(userID int64) *errors.RestError {
	if userID <= 0 {
		return errors.NewBadRequestError("Invalid user id")
	}

	err := userDAO.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}

// Search to search users having given status
func (*usersService) Search(status string) (userdto.Users, *errors.RestError) {
	if status == "" {
		return nil, errors.NewBadRequestError("Invalid status")
	}
	users, err := userDAO.FindByStatus(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error)
	}

	return users, nil
}

// Login to login user in the system
func (*usersService) Login(request userdto.LoginRequest) (*userdto.User, *errors.RestError) {
	user := &userdto.User{
		Email:    request.Email,
		Password: crypto.GetMD5(request.Password),
	}
	if err := userDAO.FindByEmailAndPassword(user); err != nil {
		return nil, err
	}
	return user, nil
}
