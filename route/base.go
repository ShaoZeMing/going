package route

import (
	"github.com/ShaoZeMing/going/controllor/base"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("v1")
	v1.GET("/ready", base.Ready)

	return router
}
