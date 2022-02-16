package router

import (
	"gin-wall/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	AdminRouter := Router.Group("admin")
	{
		AdminRouter.POST("register", api.UserRegisters) //用户注册
		AdminRouter.POST("login", api.UserLogin)        //用户登录
		AdminRouter.POST("realname", api.RealName)      //用户实名
		AdminRouter.DELETE("delete", api.Delete)        //用户注销账号
	}
}
