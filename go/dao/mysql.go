package dao

import (
  "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)
//ָ������
const DRIVER = "mysql"

var SqlSession *gorm.DB

//���ò���ӳ��ṹ��
type conf struct {
	Url string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName string `yaml:"dbname"`
	Port string `yaml:"post"`
}


//��ȡ���ò�������
func (c *conf) getConf() *conf {
	//��ȡresources/application.yaml�ļ�
	yamlFile, err := ioutil.ReadFile("resources/application.yaml")
	//�����ִ��󣬴�ӡ������ʾ
	if err != nil {
		fmt.Println(err.Error())
	}
	//����ȡ���ַ���ת���ɽṹ��conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

//��ʼ���������ݿ⣬���ɿɲ���������ɾ�Ĳ�ṹ�ı���
func InitMySql()(err error)  {
	var c conf
	//��ȡyaml���ò���
	conf:=c.getConf()
	//��yaml���ò���ƴ�ӳ��������ݿ��url
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
	//�������ݿ�
	SqlSession,err =gorm.Open(DRIVER,dsn)
	if err !=nil{
		panic(err)
	}
	//��֤���ݿ������Ƿ�ɹ������ɹ��������쳣
	return SqlSession.DB().Ping()
}

//�ر����ݿ�����
func Close()  {
	SqlSession.Close()
}


