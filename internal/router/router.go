package router

import (
	"OpenIDProvider/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	//健康检查路由
	r.GET("/health", controller.Health)
	//注册用户信息
	r.POST("/insertUser", controller.InsertUserHandle)

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