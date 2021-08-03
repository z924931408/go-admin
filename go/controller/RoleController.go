package controller

import (
	"go-admin/go/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"go-admin/go/service"
)


func CreateRole(c *gin.Context)  {
	var role entity.Role
	c.BindJSON(&role)
	err:=service.CreateRole(&role)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"success",
			"data":role,
		})
	}
}


func GetRoleList(c *gin.Context)  {
	todoList,err :=service.GetAllRole()
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"success",
			"data":todoList,
		})
	}
}

func UpdateRole(c *gin.Context)  {
	id,ok:=c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusBadRequest,gin.H{"error":"无效的id"})
	}
	role,err:= service.GetRoleById(id)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}
	c.BindJSON(&role)
	if err= service.UpdateRole(role);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,role)
	}
}



func DeleteRoleById(c *gin.Context)  {
	id ,ok:=c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusBadRequest,gin.H{"error":"无效的id"})
	}
	if err:= service.DeleteRoleById(id);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,gin.H{id:"deleted"})
	}
}