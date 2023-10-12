package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件
func TokenAuth(c *gin.Context) {
	access_token := c.Request.Header.Get("access_token")
	// 不存在access_token属性
	if access_token == "" {
		// 跳转登陆页面
		fmt.Println("不存在token")
	}

	// 缓存中access_token不存在
	if access_token != "123" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "token校验失败！",
		})
		c.AbortWithError(http.StatusInternalServerError, errors.New("token校验失败！"))
	}
	c.Set("user_name", "gy")
	c.Set("password", "gy1")
	c.Next()
}
