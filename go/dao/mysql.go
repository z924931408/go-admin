package dao

import (
  "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

var (
	DB *gorm.DB
)

type conf struct {
	DbHost string `yaml:"host"`
	DbUser string `yaml:"user"`
	DbPassWord string `yaml:"pwd"`
	DbName string `yaml:"dbname"`
	DbPort string `yaml:"post"`
}


func InitMySql()(err error)  {
	var c conf
	conf:=c.getConf()
    print(conf)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbUser,
		conf.DbPassWord,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)

	DB,err =gorm.Open("mysql",dsn)
	if err !=nil{
		panic(err)
	}
	return DB.DB().Ping()
}

func Close()  {
	DB.Close()
}


func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("resources/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
