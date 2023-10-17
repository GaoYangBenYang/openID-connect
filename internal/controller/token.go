package controller

import (
	"OpenIDProvider/internal/middleware"
	"OpenIDProvider/internal/model"
	"OpenIDProvider/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	//（1）校验RP在请求头Authorization字段通过HTTP Basic认证传入的client_id和client_secret。
	client_id, client_secret, isBasic := c.Request.BasicAuth()
	//判断authorization认证规则
	if !isBasic {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"code": http.StatusNonAuthoritativeInfo, "message": "Unauthorized 请添加Basic认证", "data": nil})
		return
	}
	// 验证缓存中是否有client_id, client_secret
	if value := middleware.GetString(middleware.OIDC + ":" + middleware.CLIENT + ":" + client_id); value == "" {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "client_id不存在", "data": nil})
		return
	} else if value != client_secret {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "client_secret不正确", "data": nil})
		return
	}

	// （2）校验RP在请求体中传入的code、grant_type、state、redirect_uri、client_id等参数。
	var postForm = struct {
		GrantType   string `json:"grant_type"`
		RedirectURI string `json:"redirect_uri"`
		Code        string `json:"code"`
		Scope       string `json:"scope"`
		Nonce       string `json:"nonce"`
	}{}
	if err := c.ShouldBindJSON(&postForm); err != nil {
		// 返回错误信息
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "token 请求参数异常", "data": err.Error()})
		return
	}
	// 对code进行验证
	if value := middleware.GetString(middleware.OIDC + ":" + middleware.CODE + ":" + client_id + postForm.RedirectURI); strings.EqualFold(value, postForm.Code) {
		//（3）如果都校验通过，则生成access token、id token并返回。
		//生成access token
		header := model.NewHeader("HS256", "JWT")
		accessTokenPayload := model.NewAccessTokenPayload("op.com", "1", client_id, "jwt001", postForm.Nonce, postForm.Scope)
		access_token, _ := utils.EncodeTheJWT(model.NewJsonWebToken(header, accessTokenPayload))
		// base64编码
		access_token2base64, _ := utils.Base64StdEncoding(access_token)

		//缓存access_token  client_id:code:
		if err := middleware.SetString(middleware.OIDC+":"+middleware.ACCESS_TOKEN+":"+client_id+postForm.Code, access_token2base64, 60*30); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "access_token缓存失败", "data": err.Error()})
			return
		}
		//生成id token
		payload := model.NewIdTokenPayload("op.com", "1", client_id, "jwt001", postForm.Nonce, postForm.Scope)

		id_token, _ := utils.EncodeTheJWT(model.NewJsonWebToken(header, payload))

		var tokenResponse = struct {
			AccessToken string `json:"access_token"`
			TokenType   string `json:"token_type"`
			// RefreshToken   string `json:"refresh_token"`
			// ExpiresIn   string `json:"expires_in"`
			IdToken string `json:"id_token"`
		}{
			AccessToken: access_token2base64,
			TokenType:   "Bearer",
			IdToken:     id_token,
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Token请求成功", "data": tokenResponse})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "invalid_request", "data": nil})
}
