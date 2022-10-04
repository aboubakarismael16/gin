package main

import (
	"gin/jwt-gin/controller"
	"gin/jwt-gin/middlewares"
	"gin/jwt-gin/model"
	"github.com/gin-gonic/gin"
)

func main() {

	model.ConnectDataBase()

	r := gin.Default()

	public := r.Group("api")

	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)

	protected := r.Group("api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/user", controller.CurrentUser)

	r.Run()

}
