package controller

import (
	"OpenIDProvider/internal/middleware"
	"OpenIDProvider/internal/model"
	"OpenIDProvider/internal/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// http://rp.com:8081/code_flow/oidc_op?state=aHR0cDovL3JwLmNvbTo4MDgx%26response_type=code%26scope=openid%20profile%26client_id=cnAuY29t%26nonce=cdYrYNLv6wBHlBmZjWxvrQmmD
// http://rp.com:8081/code_flow/oidc_op?state=aHR0cDovL3JwLmNvbTo4MDgx&response_type=code&scope=openid%20profile&client_id=cnAuY29t&nonce=cdYrYNLv6wBHlBmZjWxvrQmmD
// http://rp.com:8081/code_flow/oidc_op?state=aHR0cDovL3JwLmNvbTo4MDgx
func AccountVerify(c *gin.Context) {
	//获取参数
	account := c.PostForm("account")
	password := c.PostForm("password")
	authz_uri := c.PostForm("authz_uri")
	//根据账号验证数据库是否存在用户
	id, err := model.SelectUserByTelephoneOrEmail(account)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "account不存在", "data": nil})
		return
	}
	//根据id查询密码
	passwordSQL := model.SelectPasswordByUserID(id)
	//身份验证正确，重定向授权接口
	if strings.EqualFold(password, passwordSQL) {
		//用户身份验证成功设置cookie
		cookieKey := "account_verify"
		cookieValue := "true"
		//存储cookie
		if err := middleware.SetString(middleware.OIDC+":"+middleware.COOKIE+":"+cookieKey, cookieValue, 0); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "cookie缓存失败", "data": err.Error()})
			return
		} else {
			//正确，重定向op授权接口，并设置名为oidc的cookie
			c.SetCookie(cookieKey, cookieValue, 1800, "/", "op.com", false, true)
			c.Redirect(http.StatusSeeOther, authz_uri)
			return
		}
	}
	//对authz_uri进行转义处理 不然两次重定向后又会出现uri截断
	//func Replace(要替换的整个字符串, 要替换的字符串, 替换成什么字符串, 要替换的次数，-1，那么就会将字符串 s 中的所有的 old 替换成 new。) string
	authz_uri = strings.Replace(authz_uri, "&", "%26", -1)
	//身份验证错误，重定向登陆页面
	c.Redirect(http.StatusSeeOther, "/login?authz_uri="+authz_uri)
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
