package controller

import (
	"OpenIDProvider/src/model"
	"net/url"

	"github.com/astaxie/beego"
)
//们通过 AuthController 来实现授权端点。当 RP 发送授权请求时，我们会根据请求参数显示登录界面。当用户提交登录信息后，我们会验证用户名和密码是否正确，并将用户信息保存在 session 中。然后，我们会生成授权码，并将授权码和 state 参数添加到重定向 URL 中，最后重定向回 RP。
//授权
type AuthorizeController struct {
	beego.Controller
}

func (a *AuthorizeController) Get() {
    // 获取授权请求参数
    clientID := a.GetString("client_id")
    redirectURI := a.GetString("redirect_uri")
    // responseType := a.GetString("response_type")
    // scope := a.GetString("scope")
    // state := a.GetString("state")

    // 验证 client_id 和 redirect_uri 是否合法
    if clientID != model.OPConfig.ClientID || redirectURI != model.OPConfig.RedirectURI {
        a.Data["json"] = map[string]interface{}{
            "error": "invalid_request",
            "error_description": "Invalid client ID or redirect URI",
        }
        a.ServeJSON()
        return
    }

    // 在授权页面上显示登录界面
    a.TplName = "login.tpl"
}

func (c *AuthorizeController) Post() {
    // 获取用户提交的登录信息
    username := c.GetString("username")
    password := c.GetString("password")

    // 验证用户名和密码是否正确
    if username != "your_username" || password != "your_password" {
        c.Data["json"] = map[string]interface{}{
            "error": "access_denied",
            "error_description": "Invalid username or password",
        }
        c.ServeJSON()
        return
    }

    // 将用户信息保存在 session 中
    c.SetSession("username", username)

    // 生成授权码
    code := "your_authorization_code"

    // 构造授权码重定向 URL
    redirectURI := c.GetString("redirect_uri")
    state := c.GetString("state")
    redirectURL, err := url.Parse(redirectURI)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "invalid_request",
            "error_description": "Invalid redirect URI",
        }
        c.ServeJSON()
        return
    }
    redirectURL.Query().Set("code", code)
    redirectURL.Query().Set("state", state)
    redirectURL.RawQuery = redirectURL.Query().Encode()

    // 重定向回 RP，返回授权码和 state 参数
    c.Redirect(redirectURL.String(), 302)
}