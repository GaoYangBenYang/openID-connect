package controller

import (
	"OpenIDProvider/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {

	msg := model.NewResponse(http.StatusOK, "用户注册成功", nil)
	c.JSON(msg.Code, msg)
}
