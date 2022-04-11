package controllers

import (
	"fmt"
	"net/http"

	"github.com/coreos/go-oidc"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	clientID     = "app1"
	clientSecret = "ZXhhbXBsZS1hcHAtc2VjcmV0"
	redirectURL = "http://localhost:8081/login"
)

var endpotin = oauth2.Endpoint{
	AuthURL:  "http://localhost:8080/authorize",
	TokenURL: "http://localhost:8080/token",
}

func init(){
	ctx := context.Background()
	//"/.well-known/openid-configuration"
	provider,err := oidc.NewProvider(ctx, "http://localhost:8080")

	if err != nil {
		// 错误处理
		fmt.Println("错误信息")
	}
	// // 配置OpenID Connect Aware OAuth2客户端
	oauth2Config := oauth2.Config{
    ClientID:     clientID,
    ClientSecret: clientSecret,
    RedirectURL:  redirectURL,
    // Discovey返回OAuth2端点
	//TODO 程序异常
    Endpoint: endpotin,
    // “OpenID”是OpenID Connect流程所需的范围。
    Scopes: []string{oidc.ScopeOpenID, "profile"},
	}

	//在响应时，提供程序可用于验证 ID 令牌。
	var verifier = provider.Verifier(&oidc.Config{ClientID: clientID})
	state:="test"
	// OAuth2 重定向保持不变。
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// 验证状态和错误。

		oauth2Token, err := oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			// 错误处理
			fmt.Println("错误信息")
		}
		fmt.Printf("oauth2Token%s",oauth2Token)
		// 从OAuth2令牌中提取ID令牌
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			// 处理缺少令牌
			fmt.Println("错误信息")
		}
		fmt.Printf("rawIDToken%s",rawIDToken)
		// 解析并验证ID令牌有效载荷
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			// handle error
			fmt.Println("错误信息")
		}
		fmt.Printf("idToken%s",idToken)
		// Extract custom claims
		var claims struct {
			Email    string `json:"email"`
			Verified bool   `json:"email_verified"`
		}
		if err := idToken.Claims(&claims); err != nil {
			// handle error
			fmt.Println("错误信息")
		}
	})
}