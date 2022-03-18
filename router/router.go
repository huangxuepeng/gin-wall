package router

import (
	"gin-wall/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	TestRouter := Router.Group("test")
	{
		TestRouter.GET("tt/:id", api.Test) //测试函数
	}

	// 公共的方法
	// PublicRouter := Router.Group("pub")
	// {

	// }
	AdminRouter := Router.Group("admin")
	{
		AdminRouter.POST("getlist", api.GetList)                             //用户查询
		AdminRouter.POST("getrealname", api.GetRealNameList)                 //实名用户的列表
		AdminRouter.POST("binning", api.BinningUser)                         //拉黑用户, 对用户的进行短暂的封号
		AdminRouter.GET("change/:id/:isReal", api.PutRealName)               //更新数据
		AdminRouter.POST("authenticationuser", api.AuthenticationUser)       //对实名信息是否成功的邮箱发送, 并且完成
		AdminRouter.DELETE("deleterealname/:id", api.DeleteRealName)         //对实名不规范用户的删除(软删除)
		AdminRouter.DELETE("deleteregistername/:id", api.DeleteRegisterName) //对用户实现注销的功能(软删除)
	}
	// 论坛路由
	// ForumRouter := Router.Group("forum")
	// {

	// }
	// 动态路由
	// DYRouter := Router.Group("dynamtic")
	// {

	// }
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", api.UserRegisters) //用户注册
		UserRouter.POST("login", api.UserLogin)        //用户登录
		UserRouter.POST("realname", api.RealName)      //用户实名
		UserRouter.GET("getid/:id", api.GetID)         //实名的显示
	}

	// c端统一加密方式
	Cpassword := Router.Group("password")
	{
		Cpassword.POST("encrypt", api.CEncrypt) //对c端的密码进行数据
	}
}
