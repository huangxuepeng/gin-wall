package initialize

import (
	"gin-wall/middleware"
	"gin-wall/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors())
	ApiGroup := Router.Group("u/v1")
	router.InitUserRouter(ApiGroup)
	return Router
}
