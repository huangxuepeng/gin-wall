package main

import (
	_ "gin-wall/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Person struct {
	ID   int    `form:"id" json:"id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/test", test)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8082")
}

// @summary test测试
// @description 协同开发接口样式测试
// @accept json
// @produce json
// @success 200 {string}
// @param id body int true "下面填写用户id"
// @param name body string true "填写姓名"
// @router /test [post]
func test(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {

		c.JSON(http.StatusOK, gin.H{

			"Success": false,
			"code":    500,
			"message": "提交失败",
			"error":   err.Error(),
			"data": gin.H{
				"user": "",
			},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"code":    200,
		"message": "提交成功",
		"data": gin.H{
			"user": "巴拉巴拉",
		},
	})
}
