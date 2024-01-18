package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *GET) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Msg":     "pong",
		"Data":    empty{},
	})
}
