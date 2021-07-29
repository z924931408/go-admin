package dao

import (
  "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)
//指定驱动
const DRIVER = "mysql"

var SqlSession *gorm.DB

//配置参数映射结构体
type conf struct {
	Url string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName string `yaml:"dbname"`
	Port string `yaml:"post"`
}


//获取配置参数数据
func (c *conf) getConf() *conf {
	//读取resources/application.yaml文件
	yamlFile, err := ioutil.ReadFile("resources/application.yaml")
	//若出现错误，打印错误提示
	if err != nil {
		fmt.Println(err.Error())
	}
	//将读取的字符串转换成结构体conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

//初始化连接数据库，生成可操作基本增删改查结构的变量
func InitMySql()(err error)  {
	var c conf
	//获取yaml配置参数
	conf:=c.getConf()
	//将yaml配置参数拼接成连接数据库的url
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
	//连接数据库
	SqlSession,err =gorm.Open(DRIVER,dsn)
	if err !=nil{
		panic(err)
	}
	//验证数据库连接是否成功，若成功，则无异常
	return SqlSession.DB().Ping()
}

//关闭数据库连接
func Close()  {
	SqlSession.Close()
}


