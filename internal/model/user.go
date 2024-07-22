package model

import "errors"

var (
	// ErrInvalidPassword is an error for invalid password
	ErrInvalidPassword = errors.New("invalid password")
	// ErrUsernameNotFound is an error for username not found
	ErrUsernameNotFound = errors.New("username not found")
)

// User is a struct that represents a user for database model
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// IsPasswordMatch is a method to check if the password is matched
func (u User) IsPasswordMatch(password string) bool {
	return u.Password == password
}

// IUserRepository is an interface that represents a user repository
type IUserRepository interface {
	Create(user User) error
	Login(username string) (User, error)
	FindByUsername(username string) (User, error)
}

// IUserUsecase is an interface that represents a user usecase
type IUserUsecase interface {
	Create(user User) error
	Login(username string, password string) (User, error)
	FindByUsername(username string) (User, error)
}
