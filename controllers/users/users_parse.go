package users

import (
	"strconv"

	userdto "github.com/DeKal/bookstore_users-api/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

var (
	userParser usersDTOParserInterface = &usersDTOParser{}
)

type usersDTOParserInterface interface {
	parseUserID(context *gin.Context) (int64, *errors.RestError)
	parseUser(context *gin.Context) (*userdto.User, *errors.RestError)
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

func (*usersDTOParser) parseUser(context *gin.Context) (*userdto.User, *errors.RestError) {
	user := &userdto.User{}
	err := context.ShouldBindJSON(&user)
	if err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		return nil, restError
	}
	return user, nil
}
