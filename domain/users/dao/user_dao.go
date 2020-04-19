package userdao

import (
	"fmt"

	usersdb "github.com/DeKal/bookstore_users-api/datasources/mysql/users_db"
	userdto "github.com/DeKal/bookstore_users-api/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/logger"
	"github.com/DeKal/bookstore_users-api/utils/errors"
	mysqlutils "github.com/DeKal/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users where id=?;"
	queryUpdate           = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDelete           = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailPwd   = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

var (
	// UserDAO contains logic working directly with DB
	UserDAO userDAOInterface = &userDAO{}
)

type userDAO struct{}
type userDAOInterface interface {
	Save(*userdto.User) *errors.RestError
	Get(*userdto.User) *errors.RestError
	Update(*userdto.User) *errors.RestError
	Delete(int64) *errors.RestError
	FindByStatus(string) (userdto.Users, *errors.RestError)
	FindByEmailAndPassword(*userdto.User) *errors.RestError
}

// Save to persist User to DB
func (*userDAO) Save(user *userdto.User) *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertUser, err := stmt.Exec(
		user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
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
func (*userDAO) Get(user *userdto.User) *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error while prepare SQL statement for get user", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	err = result.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)

	if err != nil {
		return mysqlutils.HandleGetUserError(user, err)
	}

	return nil
}

// Update to update existed User from DB
func (*userDAO) Update(user *userdto.User) *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryUpdate)
	if err != nil {
		logger.Error("Error while prepare SQL statement for update user", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("Error while execute SQL update statement", err)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

// Delete to delete existed User from DB
func (*userDAO) Delete(userID int64) *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryDelete)
	if err != nil {
		logger.Error("Error while prepare SQL statement for delete user", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userID); err != nil {
		logger.Error("Error while execute SQL delete statement", err)
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

// FindByStatus find Users by status
func (*userDAO) FindByStatus(status string) (userdto.Users, *errors.RestError) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("Error while prepare SQL statement for find user by status", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error while execute SQL queryFindUserByStatus statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer rows.Close()

	users := make(userdto.Users, 0)
	for rows.Next() {
		user := userdto.User{}
		err := rows.Scan(
			&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
		if err != nil {
			logger.Error("Error while execute SQL queryFindUserByStatus statement", err)
			return nil, errors.NewInternalServerError("Database error")
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.NewInternalServerError(fmt.Sprintf("No user found with status %s", status))
	}

	return users, nil
}

// FindByEmailAndPassword to get User from DB
func (*userDAO) FindByEmailAndPassword(user *userdto.User) *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryFindByEmailPwd)
	if err != nil {
		logger.Error("Error while prepare SQL statement for get user by email and password", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, userdto.StatusActive)
	err = result.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)

	if err != nil {
		logger.Error("Error for get user by email and password", err)
		return mysqlutils.HandleLoginUserError(user, err)
	}

	return nil
}
