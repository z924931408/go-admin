package mapper

import (
	"go-admin/go/entity"
	"go-admin/go/dao"
)



func Create(user *entity.User)(err error)   {
	if err = dao.SqlSession.Create(user).Error;err!=nil{
		return err
	}
	return
}

func SelectList()(userList []*entity.User,err error){
	if err:=dao.SqlSession.Find(&userList).Error;err!=nil{
		return nil,err
	}
	return
}

func DeleteById(id string)(err error){
	err = dao.SqlSession.Where("id=?",id).Delete(&entity.User{}).Error
	return
}


func GetUserById(id string)(user *entity.User,err error)  {
	if err = dao.SqlSession.Where("id=?",id).First(user).Error;err!=nil{
		return nil,err
	}
	return
}


func UpdateUser(user * entity.User)(err error)  {
	err = dao.SqlSession.Save(user).Error
	return
}