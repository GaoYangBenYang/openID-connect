package controller

import (
	"OpenIDProvider/internal/model"
	"OpenIDProvider/internal/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserPassVerify(c *gin.Context) {

	//验证账号密码是否正确
	userName := c.PostForm("username")
	password := c.PostForm("password")

	//验证用户名是否存在
	id, err := model.SelectUserByUserName(userName)
	if err != nil {

		return
	}

	//根据id查询密码
	passwordSQL := model.SelectPasswordByUserID(id)

	//错误，继续重定向登陆页面
	if strings.EqualFold(password, passwordSQL) {
		c.Redirect(http.StatusSeeOther, "/v1/login")
	}

	c.SetCookie("oidc_login", "oidc_login", 1800, "/", "op.com", false, true)

	//正确，重定向op授权接口，并设置名为oidc的cookie
	c.Redirect(http.StatusSeeOther, "/v1/authorization?redirect_uri=http://rp.com/code_flow/oidc_op&scope=openid+profile+email+address+phone&response_type=code&nonce=cdYrYNLv6wBHlBmZjWxvrQmmD&state=DJOfvYDSDxaPzOKRoyaTaQWCoWywdeKU&client_id=EqAfEpR492It")
}

func UserInfo(c *gin.Context) {

	//（1）校验RP在请求头Authorization字段中通过Bearer关键字传入的access token。如果校验失败，返回OIDC规定的错误响应。
	access_token, isBearer := utils.BearerAuth(c.Request)

	//判断authorization认证规则
	if !isBearer {
		c.JSON(http.StatusUnauthorized, "Unauthorized 请添加Bearer认证")
		return
	}

	log.Println(access_token)

	//解析token
	jwt, err := utils.DecodeTheJWT(access_token)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	//字符串转int
	id, _ := strconv.Atoi(jwt.Payload.Sub)

	//（2）如果校验通过，返回用户详细信息。
	userInfo := model.SelectUserInfoByUserID(id)
	c.JSON(http.StatusOK, userInfo)
}
