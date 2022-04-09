package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type UserInfo struct {
	Id int
	UUID string
	UserName string
	Password string
	Email string
	MobileNumber int
	IsViolation int
}

func GetUserInfoById(id int) (*UserInfo,error) {
	sql := orm.NewOrm()
	userInfo := UserInfo{
		Id: id,
	}
	err := sql.Read(&userInfo)
	if err != nil {
		fmt.Println("查询失败",err)
		return nil, err
	}
	return &userInfo,nil
}

func GetAllUserInfo() (*[]UserInfo,error) {
	sql := orm.NewOrm()
	var userInfoSlice []UserInfo
	////返回影响行数以及错误信息
	_, err := sql.Raw("select * from user_info").QueryRows(&userInfoSlice)
	if err!=nil {
		fmt.Println("查询失败",err)
		return nil, err
	}
	return &userInfoSlice,nil
}