package ginutils

import "github.com/gin-gonic/gin"

const (
	publicHeader        = "X-Public"
	publicHeaderContent = "true"
)

// IsPublicHeader check if request header is public or not
func IsPublicHeader(context *gin.Context) bool {
	return context.GetHeader(publicHeader) == publicHeaderContent
}
