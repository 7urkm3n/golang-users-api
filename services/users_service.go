package services

import (
	"github.com/7urkm3n/bookstore_users-api/domain/users"
	"github.com/7urkm3n/bookstore_users-api/utils/errors"
)

// FindUser returns user
func FindUser() {}

// GetUser returns user
func GetUser(ID int64) (*users.User, *errors.RestError) {
	user := &users.User{ID: ID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser creates user
func CreateUser(user *users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return user, nil
}
