package app

import (
	"github.com/DeKal/bookstore_users-api/src/controllers/users"
	usersdb "github.com/DeKal/bookstore_users-api/src/datasources/mysql/users_db"
	"github.com/DeKal/bookstore_users-api/src/domain/users/dao"

	"github.com/DeKal/bookstore_users-api/src/logger"
	"github.com/DeKal/bookstore_users-api/src/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication start the application
func StartApplication() {
	usersDB := usersdb.GetNewClientConnection()
	usersDao := dao.NewUserDao(usersDB)
	usersService := services.NewUsersService(usersDao)
	usersController := users.NewController(usersService)

	mapUrls(usersController)
	logger.Info("About to start Application...")
	router.Run(":9001")
}
