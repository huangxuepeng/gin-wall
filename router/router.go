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
		AdminRouter.POST("binning", api.BinningUser)         //拉黑用户, 对用户的进行短暂的封号
		AdminRouter.PUT("change/:id/:isReal")                //更新数据

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
