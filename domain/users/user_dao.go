package users

import (
	"fmt"
	"strings"

	usersdb "github.com/DeKal/bookstore_users-api/datasources/mysql/users_db"
	"github.com/DeKal/bookstore_users-api/utils/dates"
	"github.com/DeKal/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
)

var (
	userDB = make(map[int64]*User)
)

// Save to persist User to DB
func (user *User) Save() *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dates.GetNowString()
	insertUser, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		err := fmt.Sprintf("Error while trying to save user. %s", err.Error())
		if strings.Contains(err, indexUniqueEmail) {
			err = fmt.Sprintf("User email %s has already existed", user.Email)
		}
		return errors.NewInternalServerError(err)
	}

	userID, err := insertUser.LastInsertId()
	if err != nil {
		err := fmt.Sprintf("Error while trying to save user. %s", err.Error())
		return errors.NewInternalServerError(err)
	}
	user.ID = userID
	return nil
}

// Get to get User from DB
func (user *User) Get() *errors.RestError {
	target := userDB[user.ID]
	if target == nil {
		errorMsg := fmt.Sprintf("User %d not found", user.ID)
		return errors.NewNotFoundError(errorMsg)
	}

	user.ID = target.ID
	user.Email = target.Email
	user.FirstName = target.FirstName
	user.LastName = target.LastName
	user.DateCreated = target.DateCreated

	return nil
}
