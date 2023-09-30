package controller

import (
	"OpenIDProvider/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EncodeTheJWT(c *gin.Context) {
	payload := model.NewPayload("JWT ID 1", "sessionID")
	jwt := model.NewJWT(payload)
	jwt_token := jwt.EncodeTheJWT()
	var data = struct {
		JWTToken string `json:"jwt_token"`
	}{
		JWTToken: jwt_token,
	}
	model.Msg = model.NewResponse(http.StatusOK, "JWT Token编码成功！", data)
	c.JSON(model.Msg.Code, model.Msg)
}

func DecodeTheJWT(c *gin.Context) {
	var jwt_token = struct {
		JWTToken string `json:"jwt_token"`
	}{}
	if err := c.ShouldBindJSON(&jwt_token); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, "JSON转换出错")
		return
	}
	jwt, _ := model.DecodeTheJWT(jwt_token.JWTToken)

	model.Msg = model.NewResponse(http.StatusOK, "JWT Token解码成功!", jwt)
	c.JSON(model.Msg.Code, model.Msg)
}

func VerifyTheJWT(c *gin.Context) {
	//解析
	var jwt_token = struct {
		JWTToken string `json:"jwt_token"`
	}{}
	if err := c.ShouldBindJSON(&jwt_token); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, "JSON转换出错")
		return
	}
	jwt, signatureStr := model.DecodeTheJWT(jwt_token.JWTToken)
	//验证
	flag, logStr := model.VerifyTheJWT(jwt, signatureStr)
	if flag {
		model.Msg = model.NewResponse(http.StatusOK, logStr, flag)
	} else {
		model.Msg = model.NewResponse(http.StatusCreated, logStr, flag)
	}
	c.JSON(model.Msg.Code, model.Msg)
}
