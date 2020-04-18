package users

import (
	"net/http"

	userdto "github.com/DeKal/bookstore_users-api/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/services"
	ginutils "github.com/DeKal/bookstore_users-api/utils/gin_utils"
	"github.com/gin-gonic/gin"
)

// Get getting users from bookstore
func Get(context *gin.Context) {
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
	isPublic := ginutils.IsPublicHeader(context)
	context.JSON(http.StatusOK, target.Marshall(isPublic))
}

// Create creating user for bookstore
func Create(context *gin.Context) {
	user, err := parseUser(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	target, createErr := services.CreateUser(*user)
	if createErr != nil {
		context.JSON(createErr.Status, createErr)
		return
	}
	isPublic := ginutils.IsPublicHeader(context)
	context.JSON(http.StatusCreated, target.Marshall(isPublic))
}

// Update updating user for bookstore
func Update(context *gin.Context) {
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
	isPublic := ginutils.IsPublicHeader(context)
	context.JSON(http.StatusOK, target.Marshall(isPublic))
}

// Patch updating user for bookstore
func Patch(context *gin.Context) {
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
	isPublic := ginutils.IsPublicHeader(context)
	context.JSON(http.StatusOK, target.Marshall(isPublic))
}

// Delete updating user for bookstore
func Delete(context *gin.Context) {
	userID, err := parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	deleteErr := services.DeleteUser(userID)
	if deleteErr != nil {
		context.JSON(deleteErr.Status, deleteErr)
		return
	}
	context.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search search for users for a given condition query
func Search(context *gin.Context) {
	status := context.Query("status")

	users, err := services.Search(status)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	isPublic := ginutils.IsPublicHeader(context)
	context.JSON(http.StatusOK, users.Marshall(isPublic))
}
