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
	Iss string `json:"iss"` //签发者
	Sub string `json:"sub"` //主题
	Aud string `json:"aud"` //接收者
	Nbf int    `json:"nbf"` //生效时间
	Exp int    `json:"exp"` //过期时间
	Iat int    `json:"iat"` //签发时间
	Jti string `json:"jti"` //编号

	//自定义数据自段
	Session_state string `json:"session_state"`
}

func NewPayload(iss, sub, aud string, jti, session_state string) *Payload {
	//获取当前unix时间戳
	nowUnix := int(time.Now().Unix())

	return &Payload{
		Iss:           iss,
		Sub:           sub,
		Aud:           aud,
		Nbf:           nowUnix,
		Exp:           nowUnix + 300,
		Iat:           nowUnix,
		Jti:           jti,
		Session_state: session_state,
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
