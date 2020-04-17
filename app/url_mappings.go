package app

import (
	"github.com/DeKal/bookstore_users-api/controllers/ping"
	"github.com/DeKal/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// router.GET("/users/search/:user_id", controllers.SearchUser)
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Patch)
	router.DELETE("users/:user_id", users.Delete)
}
