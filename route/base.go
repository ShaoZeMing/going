package route

import (
	"github.com/ShaoZeMing/going/controllor/base"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	v1 := r.Group("v1")
	{
		v1.GET("/ready", base.Ready)

	}
}
