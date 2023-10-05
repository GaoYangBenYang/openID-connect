package router

import (
	"OpenIDProvider/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	//静态资源
	r.LoadHTMLGlob("internal/static/*")
	v1 := r.Group("/v1")
	{
		//健康检查路由
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "openID connect provider service is working.")
		})
		//3、GET op.com/authorization：授权接口（Authorization Endpoint），需实现：
		//（1）接收并校验RP在查询参数中传入的redirect_uri、scope、response_type、client_id、state、nonce。如果校验失败，返回OIDC规定的错误响应。
		// （本例将这个带有查询参数的完整URI称为authz_uri，后面会用到它）
		//（2）检查用户是否已在OP登录（检查名为oidc的cookie）。
		// 如果未登录，则重定向到OP登录页面GET op.com/login.html?authz_uri=...（在查询参数中传入authz_uri）；
		// 如果已登录，则执行授权逻辑，将授权码等回传参数与RP提供的redirect_uri组装成完整URI，通过浏览器重定向，即返回：
		// HTTP/1.1 303 See Other
		// Location: http://rp.com/code_flow/oidc_op
		// ?state=DJOfvYDSDxaPzOKRoyaTaQWCoWywdeKU
		// &scope=openid+profile+email+address+phone
		// &code=Z0FBQUFBQmVjc2ExRTJvZmpGN1FNTERzV0NQOUVieHRHZUdLdUd5V0dvZUMzTzliS3hjeUVpVUpxN21GWWhhaTlvalVyblVOVXI2XzJVZG1vZlctNkRNZlpTanpFM2hVSzdKaWliSDUxX1RhcW5pUk9ScTRMNW1icUh6WlJGamxIWkJhbmFnakhxbUdTenJpZVowQ3dEbDh5c3Z1ZDBpOWFGSXBtMHBRSk1IUFdPTVBxUmtzUnEzVHFCWDlYMDhXUkItWXVWczkwT0Fjb1M1bktIalZCUUdSd2p3b2lnUGZYazhEODdYekhnRTJVdzZjRkR3U01SUT0%3D
		// &iss=http%3A%2F%2Fop.com
		// &client_id=EqAfEpR492It
		// 传参：
		// state：将RP传入的state原样返回。
		// code：OP签发的授权码。
		// scope：OP批准的资源权限。
		// iss：授权码的签发人，即OP的域名。
		// client_id：授权码的签发对象，即RP的client_id。
		v1.GET("/authorization", controller.Authorization)

		//登陆页面
		v1.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})

		// 4、POST op.com/user_pass/verify：验证账密接口，用户在OP登录页面输入账密，表单提交时触发此接口。其入参为账密和authz_uri，负责进行表单认证：
		//（1）如果账密错误，则仍然重定向回OP登录页面。
		//（2）如果账密正确，则构建OP自身的会话状态（设置名为oidc的cookie），通过浏览器重定向到authz_uri，即返回：
		// HTTP/1.1 303 See Other
		// Set-Cookie: oidc="1584580277|BEWREbvcxoKRTD/6bin6mA==|eeu5GgtXvmGysKMmRDWISee85yZZolGc0zBgdQ==|y/NQfIBzH01PcdK+BzBPXA=="; expires=Thu, 19-Mar-2020 01:16:17 GMT; HttpOnly; path=/
		// Location: http://op.com/authorization
		// redirect_uri=http://rp.com/code_flow/oidc_op
		// &scope=openid+profile+email+address+phone
		// &response_type=code
		// &nonce=cdYrYNLv6wBHlBmZjWxvrQmm
		// &state=DJOfvYDSDxaPzOKRoyaTaQWCoWywdeKU
		// &client_id=EqAfEpR492It
		v1.POST("/user_pass/verify", controller.UserPassVerify)

		//1、POST op.com/registration：RP向OP进行注册的接口，OP返回client_id和client_secret给RP。
		v1.POST("/registration", controller.Registration)

		// 2、GET op.com/static/jwks.json：OP提供JWT签名秘钥的接口。
		// v1.GET("/static/jwks.json", controller.Jwks)

		//令牌接口（Token Endpoint）
		// （1）校验RP在请求头Authorization字段通过HTTP Basic认证传入的client_id和client_secret。
		// （2）校验RP在请求体中传入的code、grant_type、state、redirect_uri、client_id等参数。如果校验失败，返回OIDC规定的错误响应。
		// （3）如果都校验通过，则生成access token、id token并返回。
		// HTTP/1.1 200 OK
		// Content-Type: application/json
		// {
		// 	"state": "DJOfvYDSDxaPzOKRoyaTaQWCoWywdeKU",
		// 	"scope": "openid profile email address phone",
		// 	"access_token":"Z0FBQUFBQmVjdWxLcFBWVnFzMmJ3ckcxa1ptM21kS1lDd01rY0hsVUk1WFlkUTk2MXlrN3BkdmpjTGx0VGxFeXg0SlNfVWExNldqeExRZ0RhdWt6b1Ewdy12aWlueTRMc29sc01WOHg1QXhxWjFiU1BBMmhGX2FQZnZhSUJmQVJ4eUphZ0xWRjk3M3M5MUZ0eTlVZEJxOXFtSGNuZlNaS3BqU0ZKSFBoZXE1dWdjOE13alZhLTh3Z1NVQ191V3p5a2tCSjRsT01fYVVFTV9DQmc5UFFVZnBiUmVHSXJvdzdVYzctMzJwaHVoWkFDem5MUmo4YnB2WT0=",
		// 	"token_type": "Bearer",
		// 	"id_token":"eyJhbGciOiJSUzI1NiIsImtpZCI6InZPQnl2cjBRbTBfVUI4RFZwMkRjRVNUeWt0OHZ5bV9PS1F2VjBhRm1yQWMifQ.eyJpc3MiOiAiaHR0cDovLzEyNy4wLjAuMTo4MDAyIiwgInN1YiI6ICIwMTFkYmUwMWFhYjZkZDQ1MWJlNDIzYmI0ZDI0MWFlOWM0OTA0YzJjN2FkOTk0Yzk4YmRhOGRkNmI3Y2VhYWYwIiwgImF1ZCI6IFsicTg5TFBjNkkwQTRqIl0sICJleHAiOiAxNTg0Njc1NTMwLCAiYWNyIjogInBhc3N3b3JkIiwgImlhdCI6IDE1ODQ1ODkxMzAsICJzaWQiOiAiZGNiN2E0ODlmM2JkNWY1Y2E0YmViZGFhYWY3ZjczMzY4ZjJkZTRhMDQ4YzYwNGUyMjBmOTRhNTQiLCAibm9uY2UiOiAiVzR6NldpUzF1djlwME1jMnJaNWc1Yng3In0.Jfvl4aCJy54YRbWobj7ozQUkfA2XezHtDwZhu7t5cNaguUuxNJ-epGTaub3DfmGcXI__CB_BXuQ-phWXqbz7YQ0jbwk6HtO6pGJCHfxGmcEHisM0z_-6BwJVrm6JbVw90m4zdmen5F_palkHyI4giYtrbNA8bIAraG-pZ5jZRJOmTIWHNGKopIHhUuzv39H1Ydgn5WROgz9lk24vHmyqiXiyCl2GXFcso6tEHtU9rM5oaGbIrZb6M0HfbxgmoagAw9Z9yG3p6DDihsiHUjWVccZ8o_IwS6NfJb16WFE2NoGlUBvv3Vt7VFoJJlNtTSjc7CMCij1p8k_FiN7nPMoq8w"
		// }
		v1.POST("/token", controller.Token)

		//用户详情接口（UserInfo Endpoint）
		v1.GET("/userinfo", controller.UserInfo)
	}
}
