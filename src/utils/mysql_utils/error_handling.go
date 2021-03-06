package mysqlutils

import (
	"fmt"
	"strings"

	"github.com/DeKal/bookstore_users-api/src/domain/users/dto"
	"github.com/DeKal/bookstore_utils-go/errors"
	"github.com/DeKal/bookstore_utils-go/logger"
	"github.com/go-sql-driver/mysql"
)

const (
	duplicatedKey                = 1062
	errorSaveUser                = "Error while trying to save user. %s"
	errorUserEmailExisted        = "User email %s has already existed"
	errorGetUser                 = "Error while trying to get user %d. %s"
	errorUserNotExisted          = "User %d does not exist"
	errorLoginUser               = "Something wrong when login in the user %d. %s"
	errorUserNotHavingCredential = "User %d does not have valid credential"
	errNoRow                     = "no rows in result set"
)

// HandleCommonError handle error with messages only
func HandleCommonError(logMsg string, err error) *errors.RestError {
	logger.Error(logMsg, err)
	return errors.NewInternalServerError("Database Error!!!")
}

// HandleSaveUserError handle error when saving a user
func HandleSaveUserError(user *dto.User, err error) *errors.RestError {
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
	logger.Error(errMsg, err)
	return errors.NewInternalServerError(errMsg)
}

func returnEmailExistedWhenSavingUser(email string) *errors.RestError {
	errMsg := fmt.Sprintf(errorUserEmailExisted, email)
	logger.Error(errMsg, nil)
	return errors.NewInternalServerError(errMsg)
}

// HandleGetUserError handle error when getting a user
func HandleGetUserError(user *dto.User, err error) *errors.RestError {
	errMsg := fmt.Sprintf(errorGetUser, user.ID, err.Error())
	if strings.Contains(errMsg, errNoRow) {
		errMsg = fmt.Sprintf(errorUserNotExisted, user.ID)
	}
	logger.Error(errMsg, err)
	return errors.NewInternalServerError(errMsg)
}

// HandleLoginUserError handle error when login a user
func HandleLoginUserError(user *dto.User, err error) *errors.RestError {
	logger.Error("Error for get user by email and password", err)
	errMsg := fmt.Sprintf(errorLoginUser, user.ID, err.Error())
	if strings.Contains(errMsg, errNoRow) {
		errMsg = fmt.Sprintf(errorUserNotHavingCredential, user.ID)
	}
	logger.Error(errMsg, err)
	return errors.NewInternalServerError(errMsg)
}

//HandleFindByStatusError handle error for find user by status
func HandleFindByStatusError(status string, err error) *errors.RestError {
	return errors.NewInternalServerError(fmt.Sprintf("No user found with status %s", status))
}
