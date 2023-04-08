package controller

import (
	"OpenIDProvider/src/model"

	"github.com/astaxie/beego"
)

type TokenController struct {
    beego.Controller
}

func (t *TokenController) Post() {
    // 获取授权码和 redirect_uri 参数
    grantType := t.GetString("grant_type")
    code := t.GetString("code")
    redirectURI := t.GetString("redirect_uri")

    // 验证 grant_type 参数是否正确
    if grantType != "authorization_code" {
        t.Data["json"] = map[string]interface{}{
            "error": "unsupported_grant_type",
            "error_description": "Unsupported grant type",
        }
        t.ServeJSON()
        return
    }

    // 验证 redirect_uri 参数是否正确
    if redirectURI != model.OPConfig.RedirectURI {
        t.Data["json"] = map[string]interface{}{
            "error": "invalid_request",
            "error_description": "Invalid redirect URI",
        }
        t.ServeJSON()
        return
    }

    // 验证授权码是否正确
    if code != "your_authorization_code" {
        t.Data["json"] = map[string]interface{}{
            "error": "invalid_grant",
            "error_description": "Invalid authorization code",
        }
        t.ServeJSON()
        return
    }

    // 生成访问令牌和刷新令牌
    accessToken := "your_access_token"
    refreshToken := "your_refresh_token"

    // 构造令牌响应
    tokenResponse := map[string]interface{}{
        "access_token": accessToken,
        "token_type": "Bearer",
        "expires_in": 3600,
		"refresh_token": refreshToken,
}

	// 将令牌响应转换为 JSON 格式并返回
	t.Data["json"] = tokenResponse
	t.ServeJSON()
}