package main

import (
	"OpenIDProvider/internal/config"
	"OpenIDProvider/internal/middleware/kafka"
	"OpenIDProvider/internal/middleware/mysql"
	"OpenIDProvider/internal/middleware/redis"
	"OpenIDProvider/internal/router"
	_ "OpenIDProvider/internal/model"
	"github.com/gin-gonic/gin"
)

func main() {
	//开发环境，默认开发模式
	gin.SetMode(gin.DebugMode)
	//默认返回一个Engine实例，其中已经附加了Logger和Recovery中间件。
	r := gin.Default()
	//读取配置文件
	config.InitConfig()
	//注册MySQL
	mysql.InitMySQLClient()
	//注册Redis
	redis.InitRedisClient()
	//注册Kafka
	kafka.InitKafka()
	//注册自定义中间件(访问速率限制，访问处理时间，数据校验)

	//注册路由
	router.InitRouter(r)
	//监听端口
	r.Run(config.Conf.Application.Host + ":" + config.Conf.Application.Port)
}
