package api

import (
	"fmt"
	"gin-wall/dao"
	"gin-wall/models"
	"gin-wall/util"
	"log"

	"github.com/gin-gonic/gin"
)

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
	dao.DB.Model(&user).Where("student_number = ?", stu_num).Update("is_real", 1)
	util.Success(c, 200, nil, "查询成功", map[string]interface{}{"data": users})

}
