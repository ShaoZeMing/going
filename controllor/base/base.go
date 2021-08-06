package base

import (
	"github.com/gin-gonic/gin"
)

func Ready(c *gin.Context) {

	c.JSON(200, gin.H{
		"message":     "success",
		"status_code": "200",
		"data":        nil,
	})
}
