package users

import (
	"net/http"

	"github.com/DeKal/bookstore_oauth-go/oauth"
	userdto "github.com/DeKal/bookstore_users-api/src/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/src/services"
	"github.com/DeKal/bookstore_utils-go/errors"
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
	Login(*gin.Context)
}

// Get getting users from bookstore
func (*usersController) Get(context *gin.Context) {
	if err := oauth.AuthenticateRequest(context.Request); err != nil {
		context.JSON(err.Status, err)
		return
	}

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
	if oauth.GetCallerID(context.Request) == target.ID {
		context.JSON(http.StatusOK, target.MarshallPrivate())
		return
	}
	isPublic := oauth.IsPublic(context.Request)
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
	isPublic := oauth.IsPublic(context.Request)
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
	isPublic := oauth.IsPublic(context.Request)
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
	isPublic := oauth.IsPublic(context.Request)
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
	isPublic := oauth.IsPublic(context.Request)
	context.JSON(http.StatusOK, users.Marshall(isPublic))
}

func (*usersController) Login(context *gin.Context) {
	request := userdto.LoginRequest{}
	if err := context.ShouldBindJSON(&request); err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		context.JSON(restError.Status, restError)
		return
	}

	user, err := usersService.Login(request)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	isPublic := oauth.IsPublic(context.Request)
	context.JSON(http.StatusOK, user.Marshall(isPublic))
}
