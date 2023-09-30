package controller

import (
	"OpenIDProvider/internal/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	msg := model.NewResponse(http.StatusOK, "服务正常！", nil)
	payload := model.NewPayload("JWT ID 1", "高洋")
	jwt := model.NewJWT(payload)
	str := jwt.EncodeTheJWT()
	fmt.Println(str)

	jwta,stra:=model.DecodeTheJWT(str)

	
	fmt.Println(model.VerifyTheJWT(jwta,stra))
	c.JSON(msg.Code, msg)
}
