package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApplication start the application
func StartApplication() {
	mapUrls()
	router.Run(":9001")
}
