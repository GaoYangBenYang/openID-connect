package middleware

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 声明一个全局的MySQL变量
var MysqlDB *gorm.DB
//注册MySQL
func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: Config.MySQL.UserName + ":" + Config.MySQL.Password + "@tcp(" + Config.MySQL.Address + ")/" + Config.MySQL.DBName + "?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
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
