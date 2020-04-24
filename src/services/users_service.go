package services

import (
	"github.com/DeKal/bookstore_users-api/src/domain/users/dao"
	"github.com/DeKal/bookstore_users-api/src/domain/users/dto"
	"github.com/DeKal/bookstore_utils-go/crypto"
	"github.com/DeKal/bookstore_utils-go/dates"
	"github.com/DeKal/bookstore_utils-go/errors"
)

// UsersService contains business logic for users
type UsersService struct {
	userDao dao.UserDAOInterface
}

// UsersServiceInterface is an exported interface
type UsersServiceInterface interface {
	CreateUser(dto.User) (*dto.User, *errors.RestError)
	GetUser(int64) (*dto.User, *errors.RestError)
	UpdateUser(dto.User) (*dto.User, *errors.RestError)
	PatchUser(dto.User) (*dto.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	Search(string) (dto.Users, *errors.RestError)
	Login(dto.LoginRequest) (*dto.User, *errors.RestError)
}

// NewUsersService return new UsersService
func NewUsersService(userDao dao.UserDAOInterface) UsersServiceInterface {
	return &UsersService{
		userDao: userDao,
	}
}

// CreateUser create a user in DB
func (s *UsersService) CreateUser(user dto.User) (*dto.User, *errors.RestError) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	user.Status = dto.StatusActive
	user.DateCreated = dates.GetNowDBString()
	user.Password = crypto.GetMD5(user.Password)

	err = s.userDao.Save(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser get a user in DB
func (s *UsersService) GetUser(userID int64) (*dto.User, *errors.RestError) {
	if userID <= 0 {
		return nil, errors.NewBadRequestError("Invalid user id")
	}

	target := &dto.User{
		ID: userID,
	}
	err := s.userDao.Get(target)
	if err != nil {
		return nil, err
	}

	return target, nil
}

// UpdateUser update a user in DB
func (s *UsersService) UpdateUser(user dto.User) (*dto.User, *errors.RestError) {
	existedUser, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	existedUser.FirstName = user.FirstName
	existedUser.LastName = user.LastName
	existedUser.Email = user.Email

	err = s.userDao.Update(existedUser)
	if err != nil {
		return nil, err
	}

	return existedUser, nil
}

// PatchUser update a user in DB
func (s *UsersService) PatchUser(user dto.User) (*dto.User, *errors.RestError) {
	existedUser, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if user.FirstName != "" {
		existedUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		existedUser.LastName = user.LastName
	}
	if user.Email != "" {
		existedUser.Email = user.Email
	}

	err = s.userDao.Update(existedUser)
	if err != nil {
		return nil, err
	}

	return existedUser, nil
}

// DeleteUser to Delete a user with given userId
func (s *UsersService) DeleteUser(userID int64) *errors.RestError {
	if userID <= 0 {
		return errors.NewBadRequestError("Invalid user id")
	}

	err := s.userDao.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}

// Search to search users having given status
func (s *UsersService) Search(status string) (dto.Users, *errors.RestError) {
	if status == "" {
		return nil, errors.NewBadRequestError("Invalid status")
	}
	users, err := s.userDao.FindByStatus(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error)
	}

	return users, nil
}

// Login to login user in the system
func (s *UsersService) Login(request dto.LoginRequest) (*dto.User, *errors.RestError) {
	user := &dto.User{
		Email:    request.Email,
		Password: crypto.GetMD5(request.Password),
	}
	if err := s.userDao.FindByEmailAndPassword(user); err != nil {
		return nil, err
	}
	return user, nil
}
