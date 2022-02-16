package models

//验证是否为管理员
func (u UserRegister) vifyAdmin() bool {
	return u.Role == 2
}
