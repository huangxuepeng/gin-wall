package router

import (
	"gin-wall/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	// TestRouter := Router.Group("test")
	// {
	// 	// TestRouter.GET("test", api.Test) //测试函数
	// }
	AdminRouter := Router.Group("admin")
	{
		AdminRouter.POST("getlist", api.GetList)             //用户查询
		AdminRouter.POST("getrealname", api.GetRealNameList) //实名用户的列表

	}

	// ForumRouter := Router.Group("forum")
	// {

	// }
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", api.UserRegisters) //用户注册
		UserRouter.POST("login", api.UserLogin)        //用户登录
		UserRouter.POST("realname", api.RealName)      //用户实名
	}
}
