package users

import (
	"net/http"
	"strconv"

	userdto "github.com/DeKal/bookstore_users-api/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/services"
	"github.com/DeKal/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func parseUserID(context *gin.Context) (int64, *errors.RestError) {
	userID, err := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if err != nil {
		parsedError := errors.NewBadRequestError("User id must be a number")
		return -1, parsedError
	}
	return userID, nil
}

func parseUser(context *gin.Context) (*userdto.User, *errors.RestError) {
	user := &userdto.User{}
	err := context.ShouldBindJSON(&user)
	if err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		return nil, restError
	}
	return user, nil
}

// GetUser getting users from bookstore
func GetUser(context *gin.Context) {
	userID, err := parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	target, getErr := services.GetUser(userID)
	if getErr != nil {
		context.JSON(getErr.Status, getErr)
		return
	}
	context.JSON(http.StatusCreated, target)
}

// CreateUser creating user for bookstore
func CreateUser(context *gin.Context) {
	user := userdto.User{}
	err := context.ShouldBindJSON(&user)
	if err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		context.JSON(restError.Status, restError)
		return
	}

	target, createErr := services.CreateUser(user)
	if createErr != nil {
		context.JSON(createErr.Status, createErr)
		return
	}
	context.JSON(http.StatusCreated, target)
}

// UpdateUser updating user for bookstore
func UpdateUser(context *gin.Context) {
	userID, err := parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	var user *userdto.User
	user, err = parseUser(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	user.ID = userID

	target, updateErr := services.UpdateUser(*user)
	if updateErr != nil {
		context.JSON(updateErr.Status, updateErr)
		return
	}
	context.JSON(http.StatusCreated, target)
}

// PatchUser updating user for bookstore
func PatchUser(context *gin.Context) {
	userID, err := parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	var user *userdto.User
	user, err = parseUser(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	user.ID = userID

	target, updateErr := services.PatchUser(*user)
	if updateErr != nil {
		context.JSON(updateErr.Status, updateErr)
		return
	}
	context.JSON(http.StatusCreated, target)
}
