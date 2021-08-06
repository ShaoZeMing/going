package global

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Rep struct {
	Code    int         `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"date"`
}

func Success(res *Rep, c *gin.Context) {

	if res.Code >= 400 {
		err := errors.New("name 不能为空")
		fmt.Println(err)
	} else {
		c.JSON(200, res)
	}
	c.Abort()
}

func Error(code int, message string, data interface{}, c *gin.Context) {
	c.JSON(500, gin.H{
		"message":     message,
		"status_code": code,
		"data":        data,
	})

	c.Abort()
}
