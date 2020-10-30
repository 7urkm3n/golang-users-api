package users

import (
	"fmt"

	"github.com/7urkm3n/bookstore_users-api/utils/errors"
)

// user data access object(dao)
// The point of using this only DB access

var usersDB = make(map[int64]*User)

// Get returns user
func (user *User) Get() *errors.RestError {
	res := usersDB[user.ID]
	if res == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = res.ID
	user.FirstName = res.FirstName
	user.LastName = res.LastName
	user.Email = res.Email
	user.CreateAt = res.CreateAt

	return nil
}

// Save saves user
func (user *User) Save() *errors.RestError {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s alrady registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	usersDB[user.ID] = user
	return nil
}
