package binding

//实名认证时管理员的信息的发送
type Authenticationuser struct {
	VifystudentNum string `form:"vifystudent_num" json:"vifystudent_num" binding:"required,max=9,min=9"`
	Region         string `form:"region" json:"region" binding:"required"`
}
