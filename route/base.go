package route

import (
	"github.com/ShaoZeMing/going/controllor/base"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.GET("/ready", base.Ready)
	r.GET("/healthy", base.Healthy)
	r.GET("/reload", base.Reload)

	v1 := r.Group("v1")
	{
		v1.GET("/ready", base.Ready)

	}
}
