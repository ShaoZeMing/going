package core

import (
	"fmt"
	"github.com/ShaoZeMing/going/config"
	"github.com/gin-gonic/gin"
	"log"
)

//var  Application  *App

type App struct {
	WebServer *gin.Engine
	Config    *config.Config
}

func InitApp() *App {
	app := new(App).initConfig().initWebServer()
	return app
}

func (app *App) initConfig() *App {

	app.Config = config.ConfInit()
	return app

}

// set web server
func (app *App) initWebServer() *App {
	app.WebServer = gin.Default()
	return app
}

func (app *App) ReloadConfig() *App {
	app.Config = config.ConfInit()
	return app

}

func (app *App) Use(fc ...func(app *App)) {
	for _, f := range fc {
		f(app)
	}
}

// RegisterRouter 注册路由
func (app *App) RegisterRouter(route func(*gin.Engine)) {
	route(app.WebServer)
}

func (app *App) Run() {
	appConfig := app.Config.App
	log.Printf("%s %s starting at %q\n", appConfig.AppName, appConfig.RunMode, appConfig.HttpListen)
	err := app.WebServer.Run(appConfig.HttpListen)
	if err != nil {
		fmt.Printf("Can't RunWebServer: %s\n", err.Error())
	}
}
