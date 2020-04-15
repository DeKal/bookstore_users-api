package users

import (
	"net/http"
	"strconv"

	"github.com/DeKal/bookstore_users-api/domain/users"
	"github.com/DeKal/bookstore_users-api/services"
	"github.com/DeKal/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// GetUser getting users from bookstore
func GetUser(context *gin.Context) {
	userID, err := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if err != nil {
		parsedError := errors.NewBadRequestError("User id must be a number")
		context.JSON(parsedError.Status, parsedError)
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
	user := users.User{}
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
