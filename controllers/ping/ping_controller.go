package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping returns status server <Pong>
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
