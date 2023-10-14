package controller

import (
	"OpenIDProvider/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type client struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func Registration(c *gin.Context) {
	//cnAuY29tOnJwand0
	var client *client
	err := c.BindJSON(&client)
	if err != nil {
		c.JSON(http.StatusOK, model.NewResponse(http.StatusAccepted, "RP注册失败!", err))
		return
	}
	//数据缓存
	c.JSON(http.StatusOK, model.NewResponse(http.StatusOK, "RP注册成功!", client))
}
