package dto

// LoginRequest for user
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
