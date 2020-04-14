package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser getting users from bookstore
func GetUser(context *gin.Context) {
	context.String(http.StatusNotImplemented, "Implemented me!!!")
}

// CreateUser creating user for bookstore
func CreateUser(context *gin.Context) {
	context.String(http.StatusNotImplemented, "Implemented me!!!")
}

// SearchUser finding user from bookstore
func SearchUser(context *gin.Context) {
	context.String(http.StatusNotImplemented, "Implemented me!!!")
}
