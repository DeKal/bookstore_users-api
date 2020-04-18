package mysqlutils

import (
	"fmt"
	"strings"

	userdto "github.com/DeKal/bookstore_users-api/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/logger"
	"github.com/DeKal/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	duplicatedKey         = 1062
	errorSaveUser         = "Error while trying to save user. %s"
	errorUserEmailExisted = "User email %s has already existed"
	errorGetUser          = "Error while trying to get user %d. %s"
	errorUserNotExisted   = "User %d does not exist"
	errNoRow              = "no rows in result set"
)

// HandleSaveUserError handle error when saving a user
func HandleSaveUserError(user *userdto.User, err error) *errors.RestError {
	saveError, convertErr := err.(*mysql.MySQLError)
	if convertErr {
		return returnInternalServerWhenSavingUser(err)
	}
	if saveError.Number == duplicatedKey {
		return returnEmailExistedWhenSavingUser(user.Email)
	}
	return returnInternalServerWhenSavingUser(err)
}

func returnInternalServerWhenSavingUser(err error) *errors.RestError {
	errMsg := fmt.Sprintf(errorSaveUser, err.Error())
	logger.Error(errMsg, nil)
	return errors.NewInternalServerError(errMsg)
}

func returnEmailExistedWhenSavingUser(email string) *errors.RestError {
	errMsg := fmt.Sprintf(errorUserEmailExisted, email)
	logger.Error(errMsg, nil)
	return errors.NewInternalServerError(errMsg)
}

// HandleGetUserError handle error when getting a user
func HandleGetUserError(user *userdto.User, err error) *errors.RestError {
	errMsg := fmt.Sprintf(errorGetUser, user.ID, err.Error())
	if strings.Contains(errMsg, errNoRow) {
		errMsg = fmt.Sprintf(errorUserNotExisted, user.ID)
	}
	logger.Error(errMsg, nil)
	return errors.NewInternalServerError(errMsg)
}
