package model

import (
	"OpenIDProvider/internal/middleware"

	"github.com/google/uuid"
)

/*
属性名一定要与数据库名对应  大驼峰  tag中的db可写可不写，如果使用beego的orm框架并且不想使用默认Id为主键 那必须使用tag `orm:"pk"`
*/
// 用户表
type User struct {
	UserID        int       `json:"user_id" db:"user_id"`               // 用户ID
	UserUUID      uuid.UUID `json:"user_uuid" db:"user_uuid"`           // uuid
	UserNickName  string    `json:"user_nick_name" db:"user_nick_name"` // 昵称
	UserIntro     string    `json:"user_intro" db:"user_intro"`         // 个人简介
	UserSchool    string    `json:"user_school" db:"user_school"`       // 学校
	UserAvatar    string    `json:"user_avatar" db:"user_avatar"`       // 头像图片地址
	UserName      string    `json:"user_name" db:"user_name"`           // 姓名
	UserSex       string    `json:"user_sex" db:"user_sex"`             // 性别
	UserBirthday  string    `json:"user_birthday" db:"user_birthday"`   // 出生日期
	UserIDNumber  string    `json:"user_id_number" db:"user_id_number"` // 身份证号
	UserLocation  string    `json:"user_location" db:"user_location"`   // 所在地
	UserTelephone string    `json:"user_telephone" db:"user_telephone"` // 电话号码
	UserPassword  string    `json:"user_password" db:"user_password"`   // 密码
	UserEmail     string    `json:"user_email" db:"user_email"`         // 邮箱
	UserCreate    string    `json:"user_create" db:"user_create"`       // 创建时间
	UserEdited    string    `json:"user_edited" db:"user_edited"`       // 修改时间
	UserState     string    `json:"user_state" db:"user_state"`         // 用户状态 0=正常 时间字符串=冻结结束日期
	UserDelete    int       `json:"user_delete" db:"user_delete"`       // 逻辑删除:0=未删除,1=已删除
}

func NewUser(userID int, userNickName, userIntro, userSchool, userAvatar, userName, userSex, userBirthday, userIDNumber, userLocation, userTelephone, userPassword, userEmail, userCreate, userEdited, userState string, userDelete int) *User {
	return &User{
		UserID:        userID,
		UserUUID:      uuid.New(),
		UserNickName:  userNickName,
		UserIntro:     userIntro,
		UserSchool:    userSchool,
		UserAvatar:    userAvatar,
		UserName:      userName,
		UserSex:       userSex,
		UserBirthday:  userBirthday,
		UserIDNumber:  userIDNumber,
		UserLocation:  userLocation,
		UserTelephone: userTelephone,
		UserPassword:  userPassword,
		UserEmail:     userEmail,
		UserCreate:    userCreate,
		UserEdited:    userEdited,
		UserState:     userState,
		UserDelete:    userDelete,
	}
}

// 插入单挑数据
func InsertUser(user *User) (int, error) {
	result := middleware.MysqlDB.Create(user)
	return int(result.RowsAffected), result.Error
}

// 查询数据库是否存在userName,存在返回userID
func SelectUserByUserName(userName string) (int, error) {
	var user *User
	middleware.MysqlDB.Where("user_name = ?", userName).Find(&user)
	return user.UserID, nil
}

// 根据userID查询password
func SelectPasswordByUserID(id int) string {
	var user *User
	middleware.MysqlDB.Where("user_id = ?", id).Find(&user)
	return user.UserPassword
}

// 根据id查询用户信息
func SelectUserInfoByUserID(id int) *User {
	var user *User
	middleware.MysqlDB.Where("user_id = ?", id).Find(&user)
	return user
}

// 根据用户输入的电话或者邮箱查询数据库是否存在user,存在返回userID,不存在返回error
func SelectUserByTelephoneOrEmail(account string) (int, error) {
	var user *User
	middleware.MysqlDB.Where("user_telephone = ? or user_email = ?", account, account).Find(&user)
	return user.UserID, nil
}
