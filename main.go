package main

import (
	"fmt"
	"gin-wall/dao"
	"gin-wall/initialize"
	"gin-wall/util"
)

func main() {
	if err := util.InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}
	dao.InitMysql()
	Router := initialize.Routers()
	Router.Run(":8080")
}
