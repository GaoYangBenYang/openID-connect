package controller

import (
	"OpenIDProvider/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	msg := model.NewMessage(http.StatusOK, "服务正常！", nil)

	c.JSON(msg.Code, msg)
}
