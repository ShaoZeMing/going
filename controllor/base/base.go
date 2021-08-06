package base

import (
	"github.com/ShaoZeMing/going/global"
	"github.com/gin-gonic/gin"
)

func Ready(c *gin.Context) {

	r := global.Rep{Code: 401, Message: "成功", Data: map[string]string{"user": "xiaoming", "moblie": "17778020863", "desc": "你好，你对的"}}
	global.Success(&r, c)
}
