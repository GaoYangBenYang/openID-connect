package main

import (
	"OpenIDProvider/internal/middleware"
	"OpenIDProvider/internal/router"
	"OpenIDProvider/internal/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

// gin.H{"code": http.StatusCreated, "message": "", "data": }
func main() {
	//开发环境，默认开发模式
	gin.SetMode(gin.DebugMode)
	//默认返回一个Engine实例，其中已经附加了Logger和Recovery中间件。
	r := gin.Default()
	//关闭信任任何代理
	r.SetTrustedProxies(nil)
	fmt.Println(utils.GetUUID())
	//注册自定义中间件(访问速率限制，访问处理时间，数据校验)
	//注册跨域中间件
	r.Use(middleware.Cors())
	//注册路由
	router.SingleSignOnRouter(r)
	//监听端口
	r.Run(middleware.Config.Application.Host + ":" + middleware.Config.Application.Port)
}
