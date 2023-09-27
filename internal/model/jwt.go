package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
)

// JWT的头部描述
type header struct {
	//声明加密的算法 通常直接使用 HMAC SHA256
	Alg string `json:"alg"`
	//声明类型，这里是jwt
	Typ string `json:"typ"`
}

func NewHeader(alg, typ string) *header {
	return &header{
		Alg: alg,
		Typ: typ,
	}
}

// JWT的荷载（payload）
type payload struct {
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

func NewPayload(iss, sub, aud string, nbf, exp, iat int, jti, session_state string) *payload {
	return &payload{
		Iss:           iss,
		Sub:           sub,
		Aud:           aud,
		Nbf:           nbf,
		Exp:           exp,
		Iat:           iat,
		Jti:           jti,
		Session_state: session_state,
	}
}

// JWT的签名
type signature struct {
	//对称密钥
	Secret string `json:"secret"`
}

func NewSignature(secret string) *signature {
	return &signature{
		Secret: secret,
	}
}

type JWT struct {
	Header    *header
	Payload   *payload
	Signature *signature
}

// 创建JWT对象
func NewJWT(header *header, payload *payload, signature *signature) *JWT {
	return &JWT{
		Header:    header,
		Payload:   payload,
		Signature: signature,
	}
}

// 进行JWT编码
func (jwt *JWT) EncodeTheJWT() string {
	//Base64URL算法进行header和payload转化字符串
	//将数据对象JSon化
	header, headerErr := json.Marshal(jwt.Header)
	if headerErr != nil {
		log.Println(headerErr)
	}
	headerStr := base64.URLEncoding.EncodeToString(header)

	payload, payloadErr := json.Marshal(jwt.Payload)
	if payloadErr != nil {
		log.Println(payloadErr)
	}
	payloadStr := base64.URLEncoding.EncodeToString(payload)
	//将经过base64加密后的header和payload部分进行加盐sha256加密形成signature部分
	h := hmac.New(sha256.New, []byte(jwt.Signature.Secret))
	h.Write([]byte(headerStr + "." + payloadStr))
	signatureStr := base64.URLEncoding.EncodeToString(h.Sum(nil))
	//三部分组成JWT
	token := headerStr + "." + payloadStr + "." + signatureStr

	return token
}

// 进行JWT解码
func (jwt *JWT) DecodeTheJWT() {

}

// 进行JWT校验
func (jwt *JWT) VerifyTheJWT() {

}
