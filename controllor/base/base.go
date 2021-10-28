package base

import (
	"github.com/ShaoZeMing/going/common"
	"github.com/gin-gonic/gin"
)

func Ready(c *gin.Context) {

	r := common.Rep{
		Code:    401,
		Message: "成功",
		Data:    map[string]string{"user": "xiaoming", "moblie": "17778020863", "desc": "你好，你对的"},
	}
	common.Success(c, &r)
}
