package userdto

// PublicUser contains basic information of Users
type PublicUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateCreated string `json:"date_created"`
}

// PrivateUser contains information of Users except sensitive information
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// Marshall determine what users to create based on public condition
func (users Users) Marshall(isPublic bool) []interface{} {
	results := make([]interface{}, len(users))
	for index, user := range users {
		results[index] = user.Marshall(isPublic)
	}
	return results
}

// Marshall determine what user to create based on public condition
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return user.MarshallPublic()
	}
	return user.MarshallPrivate()
}

// MarshallPublic return user as PublicUser
func (user *User) MarshallPublic() *PublicUser {
	return &PublicUser{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DateCreated: user.DateCreated,
	}
}

// MarshallPrivate return user as PrivateUser
func (user *User) MarshallPrivate() *PrivateUser {
	return &PrivateUser{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		DateCreated: user.DateCreated,
		Status:      user.Status,
	}
}
