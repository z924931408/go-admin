package main

import (
	"go-admin/go/dao"
	"go-admin/go/model"
	"go-admin/go/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)


func main()  {
	//连接数据库
	err :=dao.InitMySql()
	if err !=nil{
		panic(err)
	}
	//程序退出关闭数据库连接
	defer dao.Close()
	//绑定模型
	dao.DB.AutoMigrate(&model.User{})
	//注册路由
	r :=routes.SetRouter()
	r.Run(":8085")
}


