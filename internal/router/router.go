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
	//OP登陆接口
	
}