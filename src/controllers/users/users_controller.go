package users

import (
	"net/http"

	"github.com/DeKal/bookstore_oauth-go/oauth"
	"github.com/DeKal/bookstore_users-api/src/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/src/services"
	"github.com/DeKal/bookstore_utils-go/errors"
	"github.com/gin-gonic/gin"
)

// Controller define handlers for endpoints
type Controller struct {
	service services.UsersServiceInterface
}

// ControllerInterface interface define handlers for endpoints
type ControllerInterface interface {
	Get(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Patch(*gin.Context)
	Delete(*gin.Context)
	Search(*gin.Context)
	Login(*gin.Context)
}

// NewController return new controller
func NewController(service services.UsersServiceInterface) ControllerInterface {
	return &Controller{
		service: service,
	}
}

// Get getting users from bookstore
func (c *Controller) Get(context *gin.Context) {
	if err := oauth.AuthenticateRequest(context.Request); err != nil {
		context.JSON(err.Status, err)
		return
	}

	userID, err := userParser.parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	target, getErr := c.service.GetUser(userID)
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
func (c *Controller) Create(context *gin.Context) {
	user, err := userParser.parseUser(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	target, createErr := c.service.CreateUser(*user)
	if createErr != nil {
		context.JSON(createErr.Status, createErr)
		return
	}
	isPublic := oauth.IsPublic(context.Request)
	context.JSON(http.StatusCreated, target.Marshall(isPublic))
}

// Update updating user for bookstore
func (c *Controller) Update(context *gin.Context) {
	userID, err := userParser.parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	var user *dto.User
	user, err = userParser.parseUser(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	user.ID = userID

	target, updateErr := c.service.UpdateUser(*user)
	if updateErr != nil {
		context.JSON(updateErr.Status, updateErr)
		return
	}
	isPublic := oauth.IsPublic(context.Request)
	context.JSON(http.StatusOK, target.Marshall(isPublic))
}

// Patch updating user for bookstore
func (c *Controller) Patch(context *gin.Context) {
	userID, err := userParser.parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	var user *dto.User
	user, err = userParser.parseUser(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	user.ID = userID

	target, updateErr := c.service.PatchUser(*user)
	if updateErr != nil {
		context.JSON(updateErr.Status, updateErr)
		return
	}
	isPublic := oauth.IsPublic(context.Request)
	context.JSON(http.StatusOK, target.Marshall(isPublic))
}

// Delete updating user for bookstore
func (c *Controller) Delete(context *gin.Context) {
	userID, err := userParser.parseUserID(context)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	deleteErr := c.service.DeleteUser(userID)
	if deleteErr != nil {
		context.JSON(deleteErr.Status, deleteErr)
		return
	}
	context.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search search for users for a given condition query
func (c *Controller) Search(context *gin.Context) {
	status := context.Query("status")

	users, err := c.service.Search(status)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	isPublic := oauth.IsPublic(context.Request)
	context.JSON(http.StatusOK, users.Marshall(isPublic))
}

// Login login with user info
func (c *Controller) Login(context *gin.Context) {
	request := dto.LoginRequest{}
	if err := context.ShouldBindJSON(&request); err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		context.JSON(restError.Status, restError)
		return
	}

	user, err := c.service.Login(request)
	if err != nil {
		context.JSON(err.Status, err)
		return
	}
	isPublic := oauth.IsPublic(context.Request)
	context.JSON(http.StatusOK, user.Marshall(isPublic))
}
