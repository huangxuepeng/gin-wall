package binding

//实名认证时管理员的信息的发送
type Authenticationuser struct {
	VifystudentNum string `form:"vifystudent_num" json:"vifystudent_num" binding:"required,max=9,min=9"` //管理员的学号
	Region         string `form:"region" json:"region" binding:"required"`                               //认证的状态
	ID             int    `form:"id" json:"id" binding:"required"`
}

// 删除信息的绑定
type DeleteRealNames struct {
	ID int `uri:"id" json:"id" binding:"required"`
}

//获取前端传来的学号参数
type GetStudentNumber struct {
	StudentNumber string `form:"studentnumber" json:"studentnumber" binding:"min=0,max=9"`
}
