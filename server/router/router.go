package router

import (
	"github.com/gin-gonic/gin"
	"golang-blog/controller"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.POST("/register", controller.Register)
	e.GET("register", controller.GORegister)
	e.GET("/", controller.Index)
	e.Run()
}
