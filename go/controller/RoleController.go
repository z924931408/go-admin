package controller

import (
	"go-admin/go/entity"
	"github.com/gin-gonic/gin"
	"go-admin/go/service"
	"go-admin/go/common/rsp"
)


func CreateRole(c *gin.Context)  {
	var role entity.Role
	c.BindJSON(&role)
	err:=service.CreateRole(&role)
	if err!=nil{
		rsp.Error(c,err.Error())
	}else {
		rsp.Success(c,"新增成功",role)

	}
}


func GetRoleList(c *gin.Context)  {
	todoList,err :=service.GetAllRole()
	if err!=nil{
		rsp.Error(c,err.Error())
	}else {
		rsp.Success(c,"请求成功",todoList)
	}
}

func UpdateRole(c *gin.Context)  {
	id,ok:=c.Params.Get("id")
	if !ok{
		rsp.Error(c,"无效的id")
	}
	role,err:= service.GetRoleById(id)
	if err!=nil{
		rsp.Error(c,err.Error())
		return
	}
	c.BindJSON(&role)
	if err= service.UpdateRole(role);err!=nil{
		rsp.Error(c,err.Error())
	}else {
		rsp.Success(c,"新增成功",role)
	}
}



func DeleteRoleById(c *gin.Context)  {
	id ,ok:=c.Params.Get("id")
	if !ok{
		rsp.Error(c,"无效的id")
	}
	if err:= service.DeleteRoleById(id);err!=nil{
		rsp.Error(c,err.Error())
	}else {
		rsp.Success(c,"删除成功",id)
	}
}