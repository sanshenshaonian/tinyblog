package router

import (
	"tinyblog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/index", controller.ListUser)
	e.POST("/index", controller.RegisterUser)
	e.GET("/", controller.Index)
	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.RegisterUser)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.GET("/post_index", controller.GetPostIndex)
	e.POST("/post", controller.AddPost)
	e.GET("/post", controller.GoAddPost)
	e.GET("/post_detail", controller.PostDetail)
	e.Run(":8009")
}
