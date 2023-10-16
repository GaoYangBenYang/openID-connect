package model

import (
	"time"
)

// JWT的头部描述
type Header struct {
	//声明加密的算法 通常直接使用 HMAC SHA256
	Alg string `json:"alg"`
	//声明类型，这里是jwt
	Typ string `json:"typ"`
}

func NewHeader(alg, typ string) *Header {
	return &Header{
		Alg: alg,
		Typ: typ,
	}
}

// JWT的荷载（Payload）
type Payload struct {
	//七个官方字段
	Iss   string `json:"iss"`   //签发者
	Sub   string `json:"sub"`   //主题
	Aud   string `json:"aud"`   //接收者
	Nbf   int64  `json:"nbf"`   //生效时间
	Exp   int64  `json:"exp"`   //过期时间
	Iat   int64  `json:"iat"`   //签发时间
	Jti   string `json:"jti"`   //编号
	Nonce string `json:"nonce"` //可选，不透明字符串，用于减少重播攻击。如果存在于 ID 令牌中，客户端必须验证 随机数声明值等于 身份验证请求中发送的随机数参数的值。 如果存在于身份验证请求中，则授权服务器 必须在 具有声明值的 ID 令牌 是身份验证请求中发送的随机数值。 授权服务器不应执行其他处理 在使用的随机数值上
	//自定义数据自段

}

func NewPayload(iss, sub, aud string, jti, nonce string) *Payload {
	//获取当前unix时间戳
	nowUnix := time.Now().Unix()

	return &Payload{
		Iss:   iss,
		Sub:   sub,
		Aud:   aud,
		Nbf:   nowUnix,
		Exp:   nowUnix + 1800,
		Iat:   nowUnix,
		Jti:   jti,
		Nonce: nonce,
	}
}

type JWT struct {
	Header  *Header
	Payload *Payload
}

// 创建JWT对象
func NewJWT(header *Header, payload *Payload) *JWT {
	return &JWT{
		Header:  header,
		Payload: payload,
	}
}
