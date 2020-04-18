package app

import (
	"github.com/DeKal/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication start the application
func StartApplication() {
	mapUrls()
	logger.Info("About to start Application...")
	router.Run(":9001")
}
