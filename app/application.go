package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

// StartApplication starting App
func StartApplication() {
	routes()
	router.Run(":3000")
}
