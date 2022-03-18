package api

import (
	"fmt"
	"gin-wall/binding"
	"gin-wall/dao"
	"gin-wall/global"
	"gin-wall/models"
	"gin-wall/util"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Bin struct {
	BinTime string `form:"bin_time" json:"bin_time" binding:"required"`
}

func GetList(c *gin.Context) {
	var users []models.UserRegister
	var studentnumber binding.GetStudentNumber
	//通过学号查询
	if err := c.ShouldBind(&studentnumber); err != nil {
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
	if studentnumber.StudentNumber == "" {
		result := dao.DB.Find(&users)
		if result.Error != nil {
			log.Println("GetUserList 数据库查询失败")
			return
		}
	} else {
		dao.DB.Where("student_number LIKE ?", fmt.Sprintf("%"+studentnumber.StudentNumber+"%")).Find(&users)
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
	var user models.UserRegister
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
	timeN := "2006-01-02 15:04:05"
	timeS, _ := time.Parse(timeN, b.BinTime) //解析时间

	//更新数据库
	time := timeS.Unix()
	//将新生成的数据更新数据库
	res := dao.DB.Model(&user).Where("ID = ?", 1).Update("binning_time", time)
	if res.Error != nil {
		util.Fail(c, 402, res.Error.Error(), "拉黑失败", nil)
		return
	}
	util.Success(c, 200, nil, "更新成功", nil)

}

// 实现管理员对用户进行认证成功的提交
func PutRealName(c *gin.Context) {
	var user models.UserRegister
	var is_real int
	// 获取前端传来的参数, 对参数进行操作
	id := c.Param("id")
	isReal := c.Param("isReal")
	if isReal == "true" {
		is_real = 1
	} else {
		is_real = 0
	}
	// 更新数据
	dao.DB.Model(&user).Where("ID = ?", id).Update("IsReal", is_real)
	util.Success(c, 200, nil, "操作成功", nil)
}

// 审核信息的上传
func AuthenticationUser(c *gin.Context) {
	var user models.UserRegister
	var authen binding.Authenticationuser
	if err := c.ShouldBind(&authen); err != nil {
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
	fmt.Println(authen)
	res := dao.DB.Where("student_num = ?", authen.VifystudentNum).First(&user)
	if res.Error != nil {
		return
	}
	if user.ID == 0 {
		util.Fail(c, 402, "输入的学号有误", "", nil)
		return
	}
	// vifyUser := user.ID
	// // 将ID放入规定的地方
	// dao.DB.Model()
}

// 用户实名信息的删除
func DeleteRealName(c *gin.Context) {
	var deleteUser binding.DeleteRealNames
	var User models.UserRealname
	//完成用户的删除(软删除)
	if err := c.ShouldBindUri(&deleteUser); err != nil {
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
	res := dao.DB.Delete(&User, deleteUser.ID)
	if res.Error != nil {
		util.Fail(c, 400, res.Error.Error(), "删除失败", nil)
		return
	}
	util.Success(c, 200, nil, "成功", nil)
}

// 用户注册信息的删除
func DeleteRegisterName(c *gin.Context) {
	var deleteUser binding.DeleteRealNames
	var User models.UserRegister
	//完成用户的删除(软删除)
	if err := c.ShouldBindUri(&deleteUser); err != nil {
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
	res := dao.DB.Delete(&User, deleteUser.ID)
	if res.Error != nil {
		util.Fail(c, 400, res.Error.Error(), "删除失败", nil)
		return
	}
	util.Success(c, 200, nil, "成功", nil)
}
