package middleware

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// 日志写入文件
func init() {
	// 创建、追加、读写，777，所有权限
	f, err := os.OpenFile("logs/openID-connect-provider.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Println("日志文件创建失败！", err)
	}
	//原生log记录
	log.SetOutput(f)
	//gin http请求记录
	gin.DefaultWriter = io.MultiWriter(f)

}
