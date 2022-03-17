package dao

import (
	"gin-wall/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/gin_wall?charset=utf8mb4&parseTime=True&loc=Local"

	//全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println("连接数据库失败")
	}
	DB.AutoMigrate(
		&models.Teacher{},
		&models.UserRegister{},
		&models.UserRealname{},
		&models.OnlineLog{},
		&models.DynamicInformation{},
		&models.FatherComment{},
		&models.SonComment{},
		&models.InfoLog{},
	)

}
