package service

import (
	"go-admin/go/dao"
	"go-admin/go/entity"
)



/**
新建User信息
 */
func CreateRole(role *entity.Role)(err error)  {
	if err = dao.SqlSession.Create(role).Error;err!=nil{
		return err
	}
	return
}

/**
获取user集合
 */
func GetAllRole()(roleList []*entity.Role,err error)  {
	if err:=dao.SqlSession.Find(&roleList).Error;err!=nil{
		return nil,err
	}
	return
}

/**
根据id删除user
 */
func DeleteRoleById(id string)(err error){
	err = dao.SqlSession.Where("id=?",id).Delete(&entity.Role{}).Error
	return
}

/**
根据id查询用户User
 */
func GetRoleById(id string)(role *entity.Role,err error)  {
	if err = dao.SqlSession.Where("id=?",id).First(role).Error;err!=nil{
		return nil,err
	}
	return
}

/**
更新用户信息
 */
func UpdateRole(role * entity.Role)(err error)  {
	err = dao.SqlSession.Save(role).Error
	return
}

