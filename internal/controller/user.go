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

func AccountVerify(c *gin.Context) {
	//获取参数
	account := c.PostForm("account")
	password := c.PostForm("password")
	//根据账号验证数据库是否存在用户
	id, err := model.SelectUserByTelephoneOrEmail(account)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NewResponse(http.StatusNotFound, "account不存在!", nil))
		return
	}
	//根据id查询密码
	passwordSQL := model.SelectPasswordByUserID(id)
	//身份验证正确，重定向授权接口
	if strings.EqualFold(password, passwordSQL) {
		//用户身份验证成功设置cookie
		c.SetCookie("oidc_account_verify", "true", 60, "/", "op.com", false, true)
		//正确，重定向op授权接口，并设置名为oidc的cookie
		c.Redirect(http.StatusSeeOther, "/v1/authorize?redirect_uri=http://rp.com:8081/code_flow/oidc_op&scope=openid+profile+email+address+phone&response_type=code&nonce=cdYrYNLv6wBHlBmZjWxvrQmmD&state=DJOfvYDSDxaPzOKRoyaTaQWCoWywdeKU&client_id=EqAfEpR492It")
		return
	}
	//身份验证错误，重定向登陆页面
	c.Redirect(http.StatusSeeOther, "/login")
}

func UserInfo(c *gin.Context) {
	//（1）校验RP在请求头Authorization字段中通过Bearer关键字传入的access token。如果校验失败，返回OIDC规定的错误响应。
	access_token, isBearer := utils.BearerAuth(c.Request)
	//判断authorization认证规则
	if !isBearer {
		c.JSON(http.StatusNonAuthoritativeInfo, "Unauthorized 请添加Bearer认证")
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
