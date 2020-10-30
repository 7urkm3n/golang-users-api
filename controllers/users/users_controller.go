package users

import (
	"net/http"
	"strconv"

	"github.com/7urkm3n/bookstore_users-api/domain/users"
	"github.com/7urkm3n/bookstore_users-api/services"
	"github.com/7urkm3n/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "SEARCH: implement me")
}

// GetUser returns user
func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		badErr := errors.NewBadRequestError("invalid user id")
		c.JSON(badErr.Status, badErr)
		return
	}

	user, userErr := services.GetUser(userID)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

// CreateUser creates User
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	res, saveErr := services.CreateUser(&user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, &res)
}

func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "UpdateUser: implement me")
}
