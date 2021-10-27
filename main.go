package main

import (
	"github.com/ShaoZeMing/going/core"
	"github.com/ShaoZeMing/going/middleware"
	"github.com/ShaoZeMing/going/models/dao/mysql"
	"github.com/ShaoZeMing/going/route"
	"github.com/gin-gonic/gin"
)

func main() {
	app := core.InitApp()
	app.RegisterRouter(route.Init)
	app.Use(mysql.InitMysql)
	app.WebServer.Use(gin.CustomRecovery(middleware.Recovered))
	app.Run()
}
