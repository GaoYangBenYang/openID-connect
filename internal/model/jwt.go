package model

import (
	"OpenIDProvider/internal/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
	"time"
)

// JWT的头部描述
type header struct {
	//声明加密的算法 通常直接使用 HMAC SHA256
	Alg string `json:"alg"`
	//声明类型，这里是jwt
	Typ string `json:"typ"`
}

func NewHeader() *header {
	return &header{
		Alg: config.Conf.JsonWebToken.HeaderAlg,
		Typ: config.Conf.JsonWebToken.HeaderTyp,
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

func NewPayload(jti, session_state string) *payload {
	//获取当前unix时间戳
	nowUnix := int(time.Now().Unix())

	return &payload{
		Iss:           config.Conf.JsonWebToken.PayloadIss,
		Sub:           config.Conf.JsonWebToken.PayloadSub,
		Aud:           config.Conf.JsonWebToken.PayloadAud,
		Nbf:           nowUnix,
		Exp:           nowUnix + 300,
		Iat:           nowUnix,
		Jti:           jti,
		Session_state: session_state,
	}
}

type JWT struct {
	Header  *header
	Payload *payload
}

// 创建JWT对象
func NewJWT(payload *payload) *JWT {
	header := NewHeader()
	return &JWT{
		Header:  header,
		Payload: payload,
	}
}

// JWT编码
func (jwt *JWT) EncodeTheJWT() string {
	//Base64URL算法进行header和payload转化字符串
	//将数据对象JSon化
	header, headerErr := json.Marshal(jwt.Header)
	if headerErr != nil {
		log.Println(headerErr)
	}
	headerStr := base64.RawURLEncoding.EncodeToString(header)

	payload, payloadErr := json.Marshal(jwt.Payload)
	if payloadErr != nil {
		log.Println(payloadErr)
	}
	payloadStr := base64.RawURLEncoding.EncodeToString(payload)
	//将经过base64加密后的header和payload部分进行加盐sha256加密形成signature部分
	h := hmac.New(sha256.New, []byte(config.Conf.JsonWebToken.SignatureSecretKey))
	h.Write([]byte(headerStr + "." + payloadStr))
	signatureStr := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	//三部分组成JWT
	jwtToken := headerStr + "." + payloadStr + "." + signatureStr

	return jwtToken
}

// JWT解码
func DecodeTheJWT(jwtToken string) (*JWT, string) {
	//拆分JWTToken
	first := strings.Index(jwtToken, ".")
	second := strings.LastIndex(jwtToken, ".")

	headerStr := jwtToken[:first]
	payloadStr := jwtToken[first+1 : second]
	signatureStr := jwtToken[second+1:]

	//对第一部分进行base64解码
	headerData, headerErr := base64.RawStdEncoding.DecodeString(headerStr)
	if headerErr != nil {
		log.Println("jwt heander 解码错误", headerErr)
	}
	var header *header
	json.Unmarshal(headerData, &header)

	//对第二部分(载体)进行base64解码看token是否过期
	payloadData, payloadErr := base64.RawStdEncoding.DecodeString(payloadStr)
	if payloadErr != nil {
		log.Println("jwt payload 解码错误", headerErr)
	}
	var payload *payload
	json.Unmarshal(payloadData, &payload)

	//返回jwt对象和第三部分字符串
	jwt := NewJWT(payload)

	return jwt, signatureStr
}

// JWT校验
func VerifyTheJWT(jwt *JWT, signatureStr string) (bool, string) {
	//获取当前unix时间戳
	timeUnix := int(time.Now().Unix())
	//校验生效时间是否到达
	if timeUnix < jwt.Payload.Nbf {
		return false, "生效时间未到达"
	}
	//校验JWT有效时间是否过期
	if timeUnix > jwt.Payload.Exp {
		return false, "token已过期"
	}
	//对其进行编码获取第三段字符串
	//Equal比较加密字符串
	jwtToken := jwt.EncodeTheJWT()
	//拆分JWTToken
	second := strings.LastIndex(jwtToken, ".")

	signatureStrTemp := jwtToken[second+1:]

	return hmac.Equal([]byte(signatureStrTemp), []byte(signatureStr)), "JWT Token校验成功！"
}
