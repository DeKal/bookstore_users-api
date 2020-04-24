package dao

import (
	"database/sql"
	"fmt"

	"github.com/DeKal/bookstore_users-api/src/domain/users/dto"
	"github.com/DeKal/bookstore_users-api/src/logger"
	mysqlutils "github.com/DeKal/bookstore_users-api/src/utils/mysql_utils"
	"github.com/DeKal/bookstore_utils-go/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users where id=?;"
	queryUpdate           = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDelete           = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailPwd   = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

// UserDAO manage direct interaction with DB
type UserDAO struct {
	client *sql.DB
}

// UserDAOInterface interface manage direct interaction with DB
type UserDAOInterface interface {
	Save(*dto.User) *errors.RestError
	Get(*dto.User) *errors.RestError
	Update(*dto.User) *errors.RestError
	Delete(int64) *errors.RestError
	FindByStatus(string) (dto.Users, *errors.RestError)
	FindByEmailAndPassword(*dto.User) *errors.RestError
}

// NewUserDao return new userDao
func NewUserDao(client *sql.DB) UserDAOInterface {
	return &UserDAO{
		client: client,
	}
}

// Save to persist User to DB
func (dao *UserDAO) Save(user *dto.User) *errors.RestError {
	stmt, err := dao.client.Prepare(queryInsertUser)
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
func (dao *UserDAO) Get(user *dto.User) *errors.RestError {
	stmt, err := dao.client.Prepare(queryGetUser)
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
func (dao *UserDAO) Update(user *dto.User) *errors.RestError {
	stmt, err := dao.client.Prepare(queryUpdate)
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
func (dao *UserDAO) Delete(userID int64) *errors.RestError {
	stmt, err := dao.client.Prepare(queryDelete)
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
func (dao *UserDAO) FindByStatus(status string) (dto.Users, *errors.RestError) {
	stmt, err := dao.client.Prepare(queryFindUserByStatus)
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

	users := make(dto.Users, 0)
	for rows.Next() {
		user := dto.User{}
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
func (dao *UserDAO) FindByEmailAndPassword(user *dto.User) *errors.RestError {
	stmt, err := dao.client.Prepare(queryFindByEmailPwd)
	if err != nil {
		logger.Error("Error while prepare SQL statement for get user by email and password", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, dto.StatusActive)
	err = result.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)

	if err != nil {
		logger.Error("Error for get user by email and password", err)
		return mysqlutils.HandleLoginUserError(user, err)
	}

	return nil
}
