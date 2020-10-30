package app

import (
	"github.com/7urkm3n/bookstore_users-api/controllers/ping"
	"github.com/7urkm3n/bookstore_users-api/controllers/users"
)

func routes() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:id", users.GetUser)
	router.PUT("/users/:id", users.UpdateUser)
	router.POST("/users", users.CreateUser)
	router.GET("/users#search", users.SearchUser)
}
