package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping Controller return pong string
func Ping(context *gin.Context) {
	context.String(http.StatusOK, "pong")
}
