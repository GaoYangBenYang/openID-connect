package controller

import (
	"OpenIDProvider/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize(c *gin.Context) {
	//根据cookie名读取，读取cookie,没有携带cookie,则重定位到OP登陆页面进行登陆
	cookie, err := c.Cookie("oidc")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	fmt.Println("cookie: ", cookie)
	//获取请求参数 重组重定向URI
	redirect_uri := c.Query("redirect_uri")
	client_id := c.Query("client_id")
	response_type := c.Query("response_type")
	scope := c.Query("scope")
	state := c.Query("state")
	nonce := c.Query("nonce")

	var authz_uri string
	//根据response_type判断使用什么模式进行授权
	if response_type == "code" {
		code, err := utils.Base64RawStdEncoding("cheshicode")
		if err != nil {
			return
		}

		authz_uri = redirect_uri + "?state=" + state + "&scope=" + scope + "&client_id=" + client_id + "&code=" + code + "&nonce" + nonce
	}
	fmt.Println("authz_uri: ", authz_uri)
	c.Redirect(http.StatusFound, authz_uri)
}
