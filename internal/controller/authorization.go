package controller

import (
	"OpenIDProvider/internal/middleware"
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
	response_type := c.Query("response_type")
	scope := c.Query("scope")
	client_id := c.Query("client_id")
	state := c.Query("state")
	redirect_uri := c.Query("redirect_uri")
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
	// 缓存中获取cookie
	value := middleware.GetString(middleware.OIDC + ":" + middleware.COOKIE + ":" + "account_verify")
	if value == "" {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "不存在cookie", "data": nil})
		return
	}

	if strings.EqualFold(cookie, value) {
		//根据response_type判断使用什么模式进行授权
		if strings.EqualFold(response_type, "code") {
			/*
				code 必填。 授权服务器生成的授权码。授权码必须在发出后不久过期，以降低泄漏风险。
				建议授权码的最长有效期为 10 分钟。客户端不得多次使用授权码。如果授权码被多次使用，
				授权服务器必须拒绝该请求，并（在可能的情况下）撤销之前根据该授权码签发的所有令牌。
				授权代码与客户端标识符和重定向 URI 绑定。
			*/
			code := utils.RandomCode()
			//缓存code 十分钟过期
			if err := middleware.SetString(middleware.OIDC+":"+middleware.CODE+":"+client_id+redirect_uri, code, 60*10); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "code缓存失败", "data": err.Error()})
				return
			}
			authz_uri = redirect_uri + "?state=" + state + "&code=" + code + "&nonce=" + nonce
		}
		c.Redirect(http.StatusFound, authz_uri)
		return
	}
}
