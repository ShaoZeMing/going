package route

import (
	"github.com/ShaoZeMing/going/controllor/base"
	v1 "github.com/ShaoZeMing/going/controllor/v1"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.GET("/ready", base.Ready)
	r.GET("/healthy", base.Healthy)
	r.GET("/reload", base.Reload)

	routeV1 := r.Group("v1")
	{
		routeV1.GET("/hello", v1.Hello)

	}
}
