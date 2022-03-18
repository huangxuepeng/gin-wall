package api

import (
	"fmt"
	"gin-wall/binding"
	"gin-wall/dao"
	"gin-wall/global"
	"gin-wall/models"
	"gin-wall/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CEncrypt(c *gin.Context) {
	var user models.UserRegister
	var cencrypt binding.GetCPassword
	if err := c.ShouldBind(&cencrypt); err != nil {
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
	fmt.Println(cencrypt)
	res := dao.DB.Where("mobile=?", cencrypt.Mobile).Find(&user)
	if res.Error != nil {
		util.Fail(c, 400, res.Error.Error(), "查找失败", nil)
		return
	}
	if user.ID == 0 {
		// 该用户不存在就进行加密, 并且返回加密完成的密码
		password := dao.Mima(cencrypt.Password)
		util.Success(c, 200, nil, "加密成功", map[string]interface{}{"data": password})
	} else {
		// 用户已经存在就进行数据的校验
		if !dao.Jiemi(user.Password, cencrypt.Password) {
			util.Fail(c, 400, "密码或手机号码失败", "登录失败", nil)
			return
		}
		util.Success(c, 200, nil, "登录成功", map[string]interface{}{"data": true})
	}
}
