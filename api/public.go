package api

import (
	"fmt"
	"gin-wall/dao"
	"gin-wall/global"
	"gin-wall/middleware"
	"gin-wall/models"
	"gin-wall/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Register struct {
	NickName      string `form:"nickname" json:"nick_name" `
	Mobile        string `form:"mobile" json:"mobile" binding:"required"`
	Avatar        string `form:"avatar" json:"avatar"`
	Email         string `form:"email" json:"email" binding:"required,email"`
	Sex           int    `form:"sex" json:"sex"`
	Constellation string `form:"constellations" json:"constellations"`
	Password      string `form:"password" json:"password" binding:"required"`
	StudentNumber string `form:"student_number" json:"student_number" binding:"required"`
}

type Login struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 公共的功能
//注册
func UserRegisters(c *gin.Context) {
	var user models.UserRegister
	var register Register
	data := map[string]interface{}{}
	if err2 := c.ShouldBind(&register); err2 != nil {
		errs, ok := err2.(validator.ValidationErrors)
		if !ok {
			util.Fail(c, 400, err2.Error(), errs.Error(), data)
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
	fmt.Println(register)
	// 验证手机号码
	if ok := util.ValidateMobile(register.Mobile); !ok {
		util.Fail(c, 400, "手机号码不正确", "注册失败", data)
		return
	}
	dao.DB.Where("mobile=?").Find(&user)
	if user.ID != 0 {
		util.Fail(c, 402, "手机号码已被注册", "注册失败", data)
		return
	}
	password := dao.Mima(register.Password)
	//生成验证码
	vifyCode := util.Random()
	//完成赋值
	user.NickName = register.NickName
	user.Mobile = register.Mobile
	user.Avatar = register.Avatar
	user.Email = register.Email
	user.Sex = uint8(register.Sex)
	user.Constellation = register.Constellation
	user.EmailAuthentication = vifyCode
	user.Password = password
	user.StudentNumber = register.StudentNumber
	user.BinningTime = strconv.Itoa(int(time.Now().Unix()))
	res := dao.DB.Create(&user) //注册成功存入数据库
	if res.Error != nil {
		util.Fail(c, 400, "失败", res.Error.Error(), nil)
		return
	}
	util.Success(c, 200, "", "注册成功", map[string]interface{}{"data": user})
}

// 用户登录
func UserLogin(c *gin.Context) {
	data := map[string]interface{}{}
	var login Login
	var user models.UserRegister
	fmt.Println(c.PostForm("mobile"))
	if err := c.ShouldBind(&login); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			util.Fail(c, 400, errs.Error(), "信息输入有误1", data)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"success": false,
			"message": "输入信息有误2",
			"error":   util.RemoveTopStruct(errs.Translate(global.Trans)),
			"data":    gin.H{},
		})
		return
	}
	res := dao.DB.Where("mobile=?", login.Mobile).First(&user)
	if res.Error != nil {
		util.Fail(c, 400, "请先完成注册再来登录吧", "登录失败", data)
		return
	}
	if user.ID == 0 {
		util.Fail(c, 422, "请先完成注册再来登录吧", "登录失败", data)
		return
	}
	if ok := dao.Jiemi(user.Password, login.Password); !ok {
		util.Fail(c, 402, "用户不存在或者密码错误", "登录失败", data)
		return
	}

	//发放token
	token, err := middleware.ReleaseToken(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	res = dao.DB.Model(&user).Update("authentication_token", token)
	if res.Error != nil {
		fmt.Println("数据库更新失败")
		return
	}
	data["token"] = token
	data["vifyCode"] = user.EmailAuthentication
	data["mobile"] = user.Mobile
	util.Success(c, 200, "", "登录成功", data)
}
