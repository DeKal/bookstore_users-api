package users

import (
	"net/http"

	userdto "github.com/DeKal/bookstore_users-api/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/services"
	ginutils "github.com/DeKal/bookstore_users-api/utils/gin_utils"
	"github.com/gin-gonic/gin"
)

var (
	// UsersController for calling service to matching URL
	UsersController usersControllerInterface = &usersController{}
	usersService                             = services.UsersService
)

type usersController struct{}
type usersControllerInterface interface {
	Get(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Patch(*gin.Context)
	Delete(*gin.Context)
	Search(*gin.Context)
}

// Get getting users from bookstore
func (*usersController) Get(context *gin.Context) {
	userID, err := userParser.parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	target, getErr := usersService.GetUser(userID)
	if getErr != nil {
		context.JSON(getErr.Status, getErr)
		return
	}
	isPublic := ginutils.IsPublicHeader(context)
	context.JSON(http.StatusOK, target.Marshall(isPublic))
}

// Create creating user for bookstore
func (*usersController) Create(context *gin.Context) {
	user, err := userParser.parseUser(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	target, createErr := usersService.CreateUser(*user)
	if createErr != nil {
		context.JSON(createErr.Status, createErr)
		return
	}
	isPublic := ginutils.IsPublicHeader(context)
	context.JSON(http.StatusCreated, target.Marshall(isPublic))
}

// Update updating user for bookstore
func (*usersController) Update(context *gin.Context) {
	userID, err := userParser.parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	var user *userdto.User
	user, err = userParser.parseUser(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	user.ID = userID

	target, updateErr := usersService.UpdateUser(*user)
	if updateErr != nil {
		context.JSON(updateErr.Status, updateErr)
		return
	}
	isPublic := ginutils.IsPublicHeader(context)
	context.JSON(http.StatusOK, target.Marshall(isPublic))
}

// Patch updating user for bookstore
func (*usersController) Patch(context *gin.Context) {
	userID, err := userParser.parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	var user *userdto.User
	user, err = userParser.parseUser(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	user.ID = userID

	target, updateErr := usersService.PatchUser(*user)
	if updateErr != nil {
		context.JSON(updateErr.Status, updateErr)
		return
	}
	isPublic := ginutils.IsPublicHeader(context)
	context.JSON(http.StatusOK, target.Marshall(isPublic))
}

// Delete updating user for bookstore
func (*usersController) Delete(context *gin.Context) {
	userID, err := userParser.parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	deleteErr := usersService.DeleteUser(userID)
	if deleteErr != nil {
		context.JSON(deleteErr.Status, deleteErr)
		return
	}
	context.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search search for users for a given condition query
func (*usersController) Search(context *gin.Context) {
	status := context.Query("status")

	users, err := usersService.Search(status)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	isPublic := ginutils.IsPublicHeader(context)
	context.JSON(http.StatusOK, users.Marshall(isPublic))
}
