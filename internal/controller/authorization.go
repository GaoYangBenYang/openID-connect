package controller

import (
	"OpenIDProvider/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
URL字符转义
用其它字符替代吧，或用全角的。
+    URL 中+号表示空格                     %2B
空格 URL中的空格可以用+号或者编码           %20
/   分隔目录和子目录                       %2F
?    分隔实际的URL和参数                   %3F
%    指定特殊字符                         %25
#    表示书签                             %23
&    URL 中指定的参数间的分隔符            %26
=    URL 中指定参数的值                   %3D
*/
func Authorize(c *gin.Context) {
	//获取请求参数 重组重定向URI
	redirect_uri := c.Query("redirect_uri")
	client_id := c.Query("client_id")
	response_type := c.Query("response_type")
	scope := c.Query("scope")
	state := c.Query("state")
	nonce := c.Query("nonce")
	//TODO检验参数
	authz_uri := redirect_uri + "%3Fstate=" + state + "%26response_type=" + response_type + "%26scope=" + scope + "%26client_id=" + client_id + "%26nonce=" + nonce
	//读取cookie,没有携带cookie,或者//TODO缓存中不存在cookie 则重定向到OP登录页面GET op.com/login.html?authz_uri=...（在查询参数中传入authz_uri）
	cookie, err := c.Cookie("account_verify")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login?authz_uri="+authz_uri)
		return
	}
	// 已登录，则执行授权逻辑，将授权码等回传参数与RP提供的redirect_uri组装成完整URI，通过浏览器重定向
	if strings.EqualFold(cookie, "true") {
		//根据response_type判断使用什么模式进行授权
		if strings.EqualFold(response_type, "code") {
			code, err := utils.Base64RawStdEncoding("cheshicode")
			if err != nil {
				return
			}
			authz_uri = redirect_uri + "?state=" + state + "&scope=" + scope + "&client_id=" + client_id + "&code=" + code + "&iss=http://op.com:8000" + "&nonce=" + nonce
		}
		c.Redirect(http.StatusFound, authz_uri)
		return
	}
}
