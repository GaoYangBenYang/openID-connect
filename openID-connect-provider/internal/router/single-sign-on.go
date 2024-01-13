package router

import (
	"OpenIDProvider/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SingleSignOnRouter(r *gin.Engine) {
	//静态资源
	r.LoadHTMLGlob("internal/static/*")
	//登陆页面
	r.GET("/login", func(c *gin.Context) {
		authz_uri := c.Query("authz_uri")
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"authz_uri": authz_uri})
	})
	v1 := r.Group("/v1")
	{
		//健康检查路由
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "openID connect provider service is working.")
		})
		//授权接口（Authorize Endpoint）
		v1.GET("/authorize", controller.Authorize)
		//登陆验证接口
		v1.POST("/account_verify", controller.AccountVerify)
		//RP向OP进行注册的接口，OP返回client_id和client_secret给RP。
		v1.POST("/registration", controller.Registration)
		// OP提供JWT签名秘钥的接口。
		// v1.GET("/static/jwks.json", controller.Jwks)
		//令牌接口（Token Endpoint）
		v1.POST("/token", controller.Token)
		//用户详情接口（UserInfo Endpoint）
		v1.GET("/userinfo", controller.UserInfo)
	}
}
