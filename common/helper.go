package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Rep struct {
	Code    int         `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"date"`
}

func Success(c *gin.Context, res *Rep) {

	if res.Code >= 400 {
	} else {

	}
	c.JSON(http.StatusOK, res)

}

func Error(c *gin.Context, res *Rep) {

	c.JSON(http.StatusInternalServerError, res)
}

func JsonFail(c *gin.Context, res *Rep) {

	c.JSON(http.StatusInternalServerError, res)
}
