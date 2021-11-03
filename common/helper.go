package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Rep struct {
	Code    int         `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"date"`
}

func Success(c *gin.Context, res *Rep) {

	c.JSON(http.StatusOK, res)

}

func Error(c *gin.Context, res *Rep) {

	c.JSON(http.StatusInternalServerError, res)
}

func JsonFail(c *gin.Context, res *Rep) {

	c.JSON(http.StatusInternalServerError, res)
}

// Try 异常捕获函数
func Try(fn func(), catchFn func(err interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catchFn(err)
		}
	}()
	fn()
}

// FenToYuan 分转元格式化为字符串
func FenToYuan(x int64) string {
	var rounded float64
	rounded = float64(x) * 100000 / 100 / 100000
	formatted := fmt.Sprintf("%.2f", rounded)
	return formatted
}
