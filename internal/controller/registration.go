package controller

import (
	"OpenIDProvider/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type client struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func Registration(c *gin.Context) {
	var client *client
	//不使用BindJSON  该函数绑定错误会直接return
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "RP注册参数不正确", "data": err.Error()})
		return
	}
	//redis存储ClientId
	if err := middleware.SetString(middleware.OIDC+":"+middleware.CLIENT+":"+client.ClientId, client.ClientSecret, 0); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "client_id注册失败", "data": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "client_id注册成功", "data": client})
	}
}
