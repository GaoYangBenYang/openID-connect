package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	//获取参数
	// fmt.Println(c.Query("redirect_uri"))
	// fmt.Println(c.Query("client_id"))
	// fmt.Println(c.Query("scope"))
	// fmt.Println(c.Query("response_type"))
	// fmt.Println(c.Query("nonce"))
	// fmt.Println(c.Query("state"))
	//没有携带cookie,则重定位到登陆页面进行登陆
	// 读取cookie,根据cookie名读取
	cookie, err := c.Cookie("oidc")
	if err != nil {
		// 直接返回cookie值
		c.Redirect(http.StatusSeeOther, "/v1/login")
		return
	}
	fmt.Println(cookie)
	c.Redirect(http.StatusSeeOther, "http://rp.com:8081/code_flow/oidc_op?state=aHR0cDovL3JwLmNvbTo4MDgxLw==&scope=openid+sub&code=Z0FBQUFBQmVjc2ExRTJvZmpGN1FNTERzV0NQOUVieHRHZUdLdUd5V0dvZUMzTzliS3hjeUVpVUpxN21GWWhhaTlvalVyblVOVXI2XzJVZG1vZlctNkRNZlpTanpFM2hVSzdKaWliSDUxX1RhcW5pUk9ScTRMNW1icUh6WlJGamxIWkJhbmFnakhxbUdTenJpZVowQ3dEbDh5c3Z1ZDBpOWFGSXBtMHBRSk1IUFdPTVBxUmtzUnEzVHFCWDlYMDhXUkItWXVWczkwT0Fjb1M1bktIalZCUUdSd2p3b2lnUGZYazhEODdYekhnRTJVdzZjRkR3U01SUT0%3D&iss=http%3A%2F%2Fop.com&client_id=EqAfEpR492It")
}
