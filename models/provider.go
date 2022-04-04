package models

// provider, err := oidc.NewProvider(ctx, "http://renzheng.com")

// if err != nil {
//     // 错误处理
// }

// // 配置OpenID Connect Aware OAuth2客户端
// oauth2Config := oauth2.Config{
//     ClientID:     clientID,
//     ClientSecret: clientSecret,
//     RedirectURL:  redirectURL,

//     // Discovey返回OAuth2端点
//     Endpoint: provider.Endpoint(),

//     // “OpenID”是OpenID Connect流程所需的范围。
//     Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
// }

// func handleRedirect(w http.ResponseWriter, r *http.Request) {
//     http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
// }

// var verifier = provider.Verifier(&oidc.Config{ClientID: clientID})

// func handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
//     // 验证状态和错误。

//     oauth2Token, err := oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
//     if err != nil {
//         // 错误处理
//     }

//     // 从OAuth2令牌中提取ID令牌
//     rawIDToken, ok := oauth2Token.Extra("id_token").(string)
//     if !ok {
//         // 处理缺少令牌
//     }

//     // 解析并验证ID令牌有效载荷
//     idToken, err := verifier.Verify(ctx, rawIDToken)
//     if err != nil {
//         // handle error
//     }

//     // Extract custom claims
//     var claims struct {
//         Email    string `json:"email"`
//         Verified bool   `json:"email_verified"`
//     }
//     if err := idToken.Claims(&claims); err != nil {
//         // handle error
//     }
// }