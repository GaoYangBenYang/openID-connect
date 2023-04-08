package controller
//我们通过 UserInfoController 来实现用户信息端点。当 RP 发送用户信息请求时，我们会根据访问令牌获取用户信息，并将用户信息响应转换为 JSON 格式并返回。如果访问令牌无效，则返回相应的错误响应。
import (
	"strings"

	"github.com/astaxie/beego"
)

type UserInfoController struct {
    beego.Controller
}

func (c *UserInfoController) Get() {
    // 获取访问令牌
    accessToken := c.Ctx.Request.Header.Get("Authorization")
    if !strings.HasPrefix(accessToken, "Bearer ") {
        c.Data["json"] = map[string]interface{}{
            "error": "invalid_request",
            "error_description": "Invalid access token",
        }
        c.ServeJSON()
        return
    }
    accessToken = accessToken[7:]

    // 根据访问令牌获取用户信息
    username := c.GetSession("username")
    if username == nil || accessToken != "your_access_token" {
        c.Data["json"] = map[string]interface{}{
            "error": "invalid_token",
            "error_description": "Invalid access token",
        }
        c.ServeJSON()
        return
    }

    // 构造用户信息响应
    userInfoResponse := map[string]interface{}{
        "sub": username,
    }

    // 将用户信息响应转换为 JSON 格式并返回
    c.Data["json"] = userInfoResponse
    c.ServeJSON()
}