package router

import (
	"OpenIDProvider/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.LoadHTMLGlob("internal/static/*")
	v1 := r.Group("/v1")
	{
		//健康检查路由
		v1.GET("/health", controller.Health)
		//注册用户信息
		v1.POST("/user", controller.InsertUserHandle)
		//生成jwt token
		v1.GET("/encode_jwt", controller.EncodeTheJWT)
		//解析jwt token
		v1.POST("/decode_jwt", controller.DecodeTheJWT)
		//验证jwt token
		v1.POST("/verify_jwt", controller.VerifyTheJWT)
		//登陆页面
		v1.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})
		//登陆接口
		v1.POST("/submit-login", func(c *gin.Context) {
			c.JSON(200, "成功")
		})
	}

	// RP准备用于鉴权请求的参数
	// RP发送请求，给OP
	//1.接受RP的鉴权请求
	// OP对用户鉴权
	// OP手机用户的鉴权信息和授权信息

	// 2.OP发送授权码给RP

	// RP使用授权码向一个端点换取访问凭证。协议称之为Token端点，但没说这个端点是不是由OP提供的。不过一般来说是
	// RP收到访问凭证，包含ID Token、Access Token
	// 客户端验证ID Token，并从中提取用户的唯一标识。前面说过这是一个JWT，唯一标识就是subject identifier

}
