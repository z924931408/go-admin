package model
// 数据库表明自定义，默认为model的复数形式，比如这里默认为 users
func (Role) TableName() string {
	return "sys_role"
}

type Role struct {
	Id int `json:"id"`//编号
	Name string `json:"name"` //角色名称
	ReMark string `json:"remark"`//备注
	DelStatus int `json:"delStatus"`//是否删除 -1：已删除   0：正常
	CreateTime int64 `json:"createTime"`//创建时间
}
