package api

import (
	"gin-wall/dao"
	"gin-wall/global"
	"gin-wall/models"
	"gin-wall/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Real struct {
	StudentNumber      string `form:"student_number" json:"student_number" binding:"required"`
	RealName           string `form:"real_name" json:"real_name" binding:"required"`
	Academy            string `form:"academy" json:"academy" binding:"required"`
	Profession         string `form:"profession" json:"profession" binding:"required"`
	Age                uint8  `form:"age" json:"age" binding:"required"`
	TeacherName        string `form:"teacher_name" json:"teacher_name"`
	PhotoStudentNumber string `form:"photo_student_number" json:"photo_student_number" binding:"required"`
	PhotoRealName      string `form:"photo_real_name" json:"photo_" binding:"required"`
	PhotoAcademy       string `form:"photo_academy" json:"photo_academy"`
	PhotoProfession    string `form:"photo_profession" json:"photo_profession" binding:"required"`
	PhotoAge           uint8  `form:"photo_age" json:"photo_age" binding:"required"`
}

//用户进行实名
func RealName(c *gin.Context) {
	//管理人员进行
	//从前端获取数据, 并且与前端输入的数据与识别的图片进行比较,
	//只要有一个不一致, 直接返回失败, 信息不正确, 如果正确等待后台管理进行审核即可
	var user models.UserRealname
	var userRes models.UserRegister
	var real Real
	if err := c.ShouldBind(&real); err != nil {
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
	//验证从前端拿到的数据是否为空
	if real.StudentNumber == "" || real.Academy == "" ||
		real.Age == 0 || real.RealName == "" ||
		real.Profession == "" {
		util.Fail(c, 403, "输入的信息不能为空", "实名失败", nil)
		return
	}
	// 判断两个是否相等
	if real.StudentNumber != real.PhotoStudentNumber ||
		real.Academy != real.PhotoAcademy ||
		real.Age != real.PhotoAge ||
		real.RealName != real.PhotoRealName ||
		real.Profession != real.PhotoProfession {
		util.Fail(c, 403, "输入的信息不一致", "实名失败", nil)
		return
	}
	dao.DB.Where("student_number = ?", real.StudentNumber).Find(&userRes)
	if userRes.ID == 0 {
		util.Fail(c, 403, "您的学号尚未注册, 请完成注册再来实名吧", "实名失败", nil)
	}
	//出入数据库即可
	data := map[string]interface{}{}
	data["student_number"] = real.StudentNumber
	data["real_name"] = real.RealName
	data["academy"] = real.Academy
	data["profession"] = real.Profession
	data["teacher_name"] = real.TeacherName
	user.StudentNumber = real.StudentNumber
	user.RealName = real.RealName
	user.Academy = real.Academy
	user.Profession = real.Profession
	user.Age = real.Age
	user.TeacherName = real.TeacherName
	user.UserRegisterID = userRes.ID
	res := dao.DB.Create(&user)
	if res.Error != nil {
		util.Fail(c, 403, res.Error.Error(), "注册失败", nil)
		return
	}
	util.Success(c, 200, nil, "实名成功等待管理员审核吧", data)
}
