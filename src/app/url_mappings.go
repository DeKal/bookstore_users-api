package app

import (
	"github.com/DeKal/bookstore_users-api/src/controllers/ping"
	"github.com/DeKal/bookstore_users-api/src/controllers/users"
)

var (
	usersController = users.UsersController
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", usersController.Create)
	router.GET("/users/:user_id", usersController.Get)
	router.PUT("/users/:user_id", usersController.Update)
	router.PATCH("/users/:user_id", usersController.Patch)
	router.DELETE("/users/:user_id", usersController.Delete)
	router.POST("/users/login", usersController.Login)

	router.GET("/internal/users/search", usersController.Search)
}
