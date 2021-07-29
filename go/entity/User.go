package entity

// 数据库表明自定义，默认为model的复数形式，比如这里默认为 users
func (User) TableName() string {
	return "sys_user"
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"` //用户名
	NickName string `json:"nickName"`
	Avatar string `json:"avatar"`
	Password string `json:"password"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	DelStatus int `json:"delStatus"`
	CreateTime int64 `json:"createTime"`
}



