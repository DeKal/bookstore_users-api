package users

import (
	"strconv"

	"github.com/DeKal/bookstore_users-api/src/domain/users/dto"
	"github.com/DeKal/bookstore_utils-go/errors"
	"github.com/gin-gonic/gin"
)

var (
	userParser usersDTOParserInterface = &usersDTOParser{}
)

type usersDTOParserInterface interface {
	parseUserID(context *gin.Context) (int64, *errors.RestError)
	parseUser(context *gin.Context) (*dto.User, *errors.RestError)
}
type usersDTOParser struct{}

func (*usersDTOParser) parseUserID(context *gin.Context) (int64, *errors.RestError) {
	userID, err := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if err != nil {
		parsedError := errors.NewBadRequestError("User id must be a number")
		return -1, parsedError
	}
	return userID, nil
}

func (*usersDTOParser) parseUser(context *gin.Context) (*dto.User, *errors.RestError) {
	user := &dto.User{}
	err := context.ShouldBindJSON(&user)
	if err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		return nil, restError
	}
	return user, nil
}
