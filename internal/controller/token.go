package controller

import (
	"OpenIDProvider/internal/model"
	"OpenIDProvider/internal/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	//（1）校验RP在请求头Authorization字段通过HTTP Basic认证传入的client_id和client_secret。
	client_id, client_secret, isBasic := c.Request.BasicAuth()
	//判断authorization认证规则
	if !isBasic {
		c.JSON(http.StatusNonAuthoritativeInfo, "Unauthorized 请添加Basic认证")
		return
	}
	// 验证缓存中是否有client_id, client_secret
	log.Println(client_id)
	log.Println(client_secret)

	// （2）校验RP在请求体中传入的code、grant_type、state、redirect_uri、client_id等参数。
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
	// 对code进行验证

	//（3）如果都校验通过，则生成access token、id token并返回。
	//生成access token
	access_token, _ := utils.Base64RawStdEncoding("ceshiaccess_token")
	//生成id token
	header := model.NewHeader("HS256", "JWT")
	payload := model.NewPayload("op.com", "5003", "rp.com", "jwt001", "")
	id_token, _ := utils.EncodeTheJWT(model.NewJWT(header, payload))
	fmt.Println("idtoken: ", id_token)
	fmt.Println(utils.Base64RawURLEncoding(id_token))
	var access_id_token = struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		IdToken     string `json:"id_token"`
		State       string `json:"state"`
	}{
		AccessToken: access_token,
		TokenType:   "Bearer",
		IdToken:     id_token,
		State:       postJson.State,
	}
	msg := model.NewResponse(http.StatusOK, "JWT Token解码成功!", access_id_token)
	c.JSON(msg.Code, msg)
}
