package main

import (
	"OpenIDProvider/internal/config"
	"OpenIDProvider/internal/middleware/mysql"
	"OpenIDProvider/internal/middleware/redis"
	"OpenIDProvider/internal/router"
	"log"
	"net/http"
	"time"
)

func init() {
	//读取配置文件
	config.InitConfig()
	//注册MySQL
	mysql.InitMySQLClient()
	//注册Redis
	redis.InitRedisClient()
	//注册Kafka
	
	//注册中间件

	//注册路由
	router.InitRouter()
}

func main() {
	//监听端口
	go func() {
		for {
			time.Sleep(time.Second)
			log.Println("进行健康检查...")
			resp, err := http.Get("http://localhost:8000/health")
			if err != nil {
				log.Println("Failed:", err)
				continue
			}
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				log.Println("Not OK:", resp.StatusCode)
				continue
			}
			break
		}
		log.Println(config.Conf.Application.Name + "启动成功！正在监听" + config.Conf.Application.Port + "端口。")
	}()
	err := http.ListenAndServe(config.Conf.Application.Host+":"+config.Conf.Application.Port, nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
