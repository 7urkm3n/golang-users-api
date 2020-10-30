package users

// user data transfer object (dto)

import (
	"strings"

	"github.com/7urkm3n/bookstore_users-api/utils/errors"
)

// User is struct
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreateAt  string `json:"created_at"`
}

// Validate returns error if not valid
func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
