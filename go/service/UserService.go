package service

import (
	"go-admin/go/dao"
	"go-admin/go/model"
)



/**
新建User信息
 */
func CreateUser(user *model.User)(err error)  {
	if err = dao.DB.Create(user).Error;err!=nil{
		return err
	}
	return
}

/**
获取user集合
 */
func GetAllUser()(userList []*model.User,err error)  {
	if err:=dao.DB.Find(&userList).Error;err!=nil{
		return nil,err
	}
	return
}

/**
根据id删除user
 */
func DeleteUserById(id string)(err error){
	err = dao.DB.Where("id=?",id).Delete(&model.User{}).Error
	return
}

/**
根据id查询用户User
 */
func GetUserById(id string)(user *model.User,err error)  {
	if err = dao.DB.Where("id=?",id).First(user).Error;err!=nil{
		return nil,err
	}
	return
}

/**
更新用户信息
 */
func UpdateUser(user * model.User)(err error)  {
	err = dao.DB.Save(user).Error
	return
}

