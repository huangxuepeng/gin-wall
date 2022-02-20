package dao

import "golang.org/x/crypto/bcrypt"

func Mima(password string) string {
	return mima(password)
}

//str是数据库中的, password是新输入的密码
func Jiemi(str string, password string) bool { //传来的数据是从数据库中拿到的
	err := bcrypt.CompareHashAndPassword([]byte(str), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func mima(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密
	if err != nil {
		panic(err)
	}
	return string(hash)

}
