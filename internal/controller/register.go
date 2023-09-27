package controller

import (
	"OpenIDProvider/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertUserHandle(c *gin.Context) {
	//Json数据绑定
	var user *model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, "JSON转换出错")
		return
	}
	_, err := model.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "用户注册失败")
		return
	}
	msg:=model.NewMessage(http.StatusOK,"用户注册成功",nil)
	c.JSON(msg.Code, msg)
}
