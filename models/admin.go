package models

//验证是否为管理员
func (u UserRegister) vifyAdmin() bool {
	return u.Role == 2
}

// 验证是否已经实名
func (u UserRegister) vifyUser() bool {
	return u.ISReal == 1
}
