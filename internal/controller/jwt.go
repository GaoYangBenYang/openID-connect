package controller

import (
	"OpenIDProvider/internal/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	//（1）校验RP在请求头Authorization字段通过HTTP Basic认证传入的client_id和client_secret。
	/*
		获取请求头Authorization字段
		解析authorization
	*/
	client_id, client_secret, isBasic := c.Request.BasicAuth()
	//判断authorization认证规则
	if !isBasic {
		c.JSON(http.StatusUnauthorized, "Unauthorized 请添加Basic认证")
		return
	}
	// （2）校验RP在请求体中传入的code、grant_type、state、redirect_uri、client_id等参数。
	// 如果校验失败，返回OIDC规定的错误响应。
	log.Println(client_id)
	log.Println(client_secret)
	var postJson = struct {
		GrantType   string `json:"grant_type"`
		RedirectURI string `json:"redirect_uri"`
		ClientId    string `json:"client_id"`
		State       string `json:"state"`
		Code        string `json:"code"`
	}{}

	if err := c.BindJSON(&postJson); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(postJson)
	//（3）如果都校验通过，则生成access token、id token并返回。

	var access_id_token = struct {
		State       string `json:"state"`
		Scope       string `json:"scope"`
		TokenType   string `json:"access_token"`
		AccessToken string `json:"token_type"`
		IdToken     string `json:"id_token"`
	}{
		State:       "DJOfvYDSDxaPzOKRoyaTaQWCoWywdeKU",
		Scope:       "openid profile email address phone",
		AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3d3cuYmVqc29uLmNvbSIsInN1YiI6IjUwMDMiLCJhdWQiOiJvbGFmIiwiaWF0IjoxNjk2MzIyODg5LCJleHAiOjE2OTYzMjM0ODl9.nXw6nSIkbzecAHfz3lDtUIOEN82L3-pEqB5IzuZm_Ag",
		TokenType:   "Bearer",
		IdToken:     "eyJhbGciOiJSUzI1NiIsImtpZCI6InZPQnl2cjBRbTBfVUI4RFZwMkRjRVNUeWt0OHZ5bV9PS1F2VjBhRm1yQWMifQ.eyJpc3MiOiAiaHR0cDovLzEyNy4wLjAuMTo4MDAyIiwgInN1YiI6ICIwMTFkYmUwMWFhYjZkZDQ1MWJlNDIzYmI0ZDI0MWFlOWM0OTA0YzJjN2FkOTk0Yzk4YmRhOGRkNmI3Y2VhYWYwIiwgImF1ZCI6IFsicTg5TFBjNkkwQTRqIl0sICJleHAiOiAxNTg0Njc1NTMwLCAiYWNyIjogInBhc3N3b3JkIiwgImlhdCI6IDE1ODQ1ODkxMzAsICJzaWQiOiAiZGNiN2E0ODlmM2JkNWY1Y2E0YmViZGFhYWY3ZjczMzY4ZjJkZTRhMDQ4YzYwNGUyMjBmOTRhNTQiLCAibm9uY2UiOiAiVzR6NldpUzF1djlwME1jMnJaNWc1Yng3In0.Jfvl4aCJy54YRbWobj7ozQUkfA2XezHtDwZhu7t5cNaguUuxNJ-epGTaub3DfmGcXI__CB_BXuQ-phWXqbz7YQ0jbwk6HtO6pGJCHfxGmcEHisM0z_-6BwJVrm6JbVw90m4zdmen5F_palkHyI4giYtrbNA8bIAraG-pZ5jZRJOmTIWHNGKopIHhUuzv39H1Ydgn5WROgz9lk24vHmyqiXiyCl2GXFcso6tEHtU9rM5oaGbIrZb6M0HfbxgmoagAw9Z9yG3p6DDihsiHUjWVccZ8o_IwS6NfJb16WFE2NoGlUBvv3Vt7VFoJJlNtTSjc7CMCij1p8k_FiN7nPMoq8w",
	}
	model.Msg = model.NewResponse(http.StatusOK, "JWT Token解码成功!", access_id_token)
	c.JSON(model.Msg.Code, model.Msg)
}
