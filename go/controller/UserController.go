package controller

import (
	"go-admin/go/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"go-admin/go/service"
)


func CreateUser(c *gin.Context)  {
	var user model.User
	c.BindJSON(&user)
	err:=service.CreateUser(&user)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"success",
			"data":user,
		})
	}
}


func GetUserList(c *gin.Context)  {
	todoList,err :=service.GetAllUser()
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

func UpdateUser(c *gin.Context)  {
	id,ok:=c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusBadRequest,gin.H{"error":"无效的id"})
	}
	user,err:= service.GetUserById(id)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}
	c.BindJSON(&user)
	if err= service.UpdateUser(user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,user)
	}
}



func DeleteUserById(c *gin.Context)  {
	id ,ok:=c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusBadRequest,gin.H{"error":"无效的id"})
	}
	if err:= service.DeleteUserById(id);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,gin.H{id:"deleted"})
	}
}