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
	client := client{
		ClientId:     "rp.com",
		ClientSecret: "rpjwt",
	}
	msg := model.NewResponse(http.StatusOK, "客户端注册成功!", client)
	c.JSON(msg.Code, msg)
}
