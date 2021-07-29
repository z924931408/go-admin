package main

import (
	"go-admin/go/dao"
	"go-admin/go/entity"
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
	dao.SqlSession.AutoMigrate(&entity.User{})
	//注册路由
	r :=routes.SetRouter()
	//启动端口为8085的项目
	r.Run(":8081")
}


