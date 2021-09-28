package controller

import (
	"go-admin/go/entity"
	"github.com/gin-gonic/gin"
	"go-admin/go/service"
	"go-admin/go/common/rsp"
)


func CreateUser(c *gin.Context)  {
	//定义一个User变量
	var user entity.User
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err:=service.CreateUser(&user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err!=nil{
		rsp.Error(c,err.Error())
	}else {
		rsp.Success(c,"新增成功",user)
	}
}


func GetUserList(c *gin.Context)  {
	todoList,err :=service.GetAllUser()
	if err!=nil{
		rsp.Error(c,err.Error())
	}else {
		rsp.Success(c,"请求成功",todoList)
	}
}

func UpdateUser(c *gin.Context)  {
	id,ok:=c.Params.Get("id")
	if !ok{
		rsp.Error(c,"无效的id")
	}
	user,err:= service.GetUserById(id)
	if err!=nil{
		rsp.Error(c,err.Error())
		return
	}
	c.BindJSON(&user)
	if err= service.UpdateUser(user);err!=nil{
		rsp.Error(c,err.Error())
	}else {
		rsp.Success(c,"更新成功",user)
	}
}



func DeleteUserById(c *gin.Context)  {
	id ,ok:=c.Params.Get("id")
	if !ok{
		rsp.Error(c,"无效的id")
	}
	if err:= service.DeleteUserById(id);err!=nil{
		rsp.Error(c,err.Error())
	}else {
		rsp.Success(c,"删除成功",id)
	}
}