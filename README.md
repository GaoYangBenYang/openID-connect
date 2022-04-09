# OpenID_Connect_OP

# SingleSignOn

# OIDC授权码模式
![OIDC授权码模式](conf\images\OIDC授权码模式.jpg)


1.授权接口
    OIDC Provider 会「检查当前用户在 OIDC Provider 的登录状态」，
    如果是未登录状态，OIDC Provider 会弹出一个登录框，与终端用户确认身份，
    登录成功后会将一个「临时授权码」（一个随机字符串）发到你的应用（「业务回调地址」）；
    如果是已登录状态，OIDC Provider 会将浏览器直接重定向到你的应用（「业务回调地址」），
    并携带「临时授权码」（一个随机字符串）。如第二、三步所示。
2.token 接口
3.用户信息接口