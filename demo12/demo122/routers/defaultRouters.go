package routers

import (
	"gin/demo12/demo122/controllers/itying"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)
		defaultRouters.GET("/news", itying.DefaultController{}.News)
		defaultRouters.GET("/shop", itying.DefaultController{}.Shop)
		defaultRouters.GET("/deleteCookie", itying.DefaultController{}.DeleteCookie)

	}
}
