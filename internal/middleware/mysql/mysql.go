package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 声明一个全局的MySQL变量
var MysqlDB *gorm.DB

func InitMySQLClient() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/codefixer?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	}), &gorm.Config{})
	if err != nil {
		log.Println("MySQL连接失败!", err)
	} else {
		MysqlDB = db
		log.Println("MySQL连接成功!")
	}

}

// 原生SQL语句
const (
	//查询所有用户信息
	SelectAllUserInfoSQL = "select * from profile"
	//根据电话号码查询用户信息
	SelectUserInfoByPhoneNumberSQL = ""
)
