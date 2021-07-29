package controller

import (
	"go-admin/go/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"go-admin/go/service"
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