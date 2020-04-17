package userdao

import (
	usersdb "github.com/DeKal/bookstore_users-api/datasources/mysql/users_db"
	userdto "github.com/DeKal/bookstore_users-api/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/utils/dates"
	"github.com/DeKal/bookstore_users-api/utils/errors"
	mysqlutils "github.com/DeKal/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users where id=?;"
	queryUpdate     = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDelete     = "DELETE FROM users WHERE id=?;"
)

// Save to persist User to DB
func Save(user *userdto.User) *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dates.GetNowString()
	insertUser, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	userID, err := insertUser.LastInsertId()
	if err != nil {
		return mysqlutils.HandleSaveUserError(user, err)
	}

	user.ID = userID
	return nil
}

// Get to get User from DB
func Get(user *userdto.User) *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	err = result.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)

	if err != nil {
		return mysqlutils.HandleGetUserError(user, err)
	}

	return nil
}

// Update to update existed User from DB
func Update(user *userdto.User) *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryUpdate)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

// Delete to delete existed User from DB
func Delete(userID int64) *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryDelete)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userID); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
