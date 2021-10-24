package main

import (
	"github.com/ShaoZeMing/going/middleware"
	"github.com/ShaoZeMing/going/route"
	"github.com/gin-gonic/gin"
)

func main() {

	r := route.Init()
	r.Use(gin.CustomRecovery(middleware.Recovered))
	err := r.Run("0.0.0.0:8011")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
