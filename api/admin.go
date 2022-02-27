package api

import (
	"fmt"
	"gin-wall/dao"
	"gin-wall/global"
	"gin-wall/models"
	"gin-wall/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Bin struct {
	BinTime string `form:"bin_time" json:"bin_time" binding:"required"`
}

func GetList(c *gin.Context) {
	var users []models.UserRegister
	//通过学号查询
	like := c.Query("query")
	names := c.PostForm("name")
	fmt.Println(names)
	if like == "" {
		result := dao.DB.Find(&users)
		if result.Error != nil {
			log.Println("GetUserList 数据库查询失败")
			return
		}
	} else {
		dao.DB.Where("student_number LIKE ?", fmt.Sprintf("%"+like+"%")).Find(&users)
	}
	util.Success(c, 200, nil, "查询成功", map[string]interface{}{"data": users})
}

func GetRealNameList(c *gin.Context) {
	var users []models.UserRealname
	var user models.UserRegister
	//允许通过学号进行查询
	dao.DB.Find(&users)
	//拿到管理员已经确认的实名信息
	stu_num := c.Query("query")
	res := dao.DB.Model(&user).Where("student_number = ?", stu_num).Update("is_real", 1)
	if res.Error != nil {
		util.Fail(c, 403, res.Error.Error(), "查询失败", nil)
		return
	}
	util.Success(c, 200, nil, "查询成功", map[string]interface{}{"data": users})
}

// 管理员对用户进行拉黑处理
func BinningUser(c *gin.Context) {
	var b Bin
	if err := c.ShouldBind(&b); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			util.Fail(c, 400, errs.Error(), "信息输入有误", nil)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"success": false,
			"message": "输入信息有误",
			"error":   util.RemoveTopStruct(errs.Translate(global.Trans)),
			"data":    gin.H{},
		})
		return
	}
	//将前端传来的时间进行格式化
	// timeN := "2006-01-02 15:04:05"
	// timeS, _ := time.Parse(timeN, b.BinTime)

}
