package base

import (
	"github.com/ShaoZeMing/going/common"
	"github.com/gin-gonic/gin"
)

//Ready 就绪
func Ready(c *gin.Context) {

	r := common.Rep{
		Code:    200,
		Message: "成功",
		Data:    map[string]string{"user": "xiaoming", "moblie": "1777***7676", "desc": ""},
	}
	common.Success(c, &r)
}

//Healthy 健康检测
func Healthy(c *gin.Context) {

	r := common.Rep{
		Code:    200,
		Message: "success",
		Data:    "",
	}
	common.Success(c, &r)
}

//Reload 配置重载接口
func Reload(c *gin.Context) {

	r := common.Rep{
		Code:    200,
		Message: "success",
		Data:    "",
	}
	common.Success(c, &r)
}
