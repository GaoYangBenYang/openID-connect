package utils

import (
	"OpenIDProvider/internal/config"
	"OpenIDProvider/internal/model"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
	"time"
)

// JWT编码
func EncodeTheJWT(jwt *model.JWT) string {
	//Base64URL算法进行header和payload转化字符串
	//将数据对象JSon化
	headerStr, headerErr := Base64RawStdEncoding(jwt.Header)
	if headerErr != nil {
		log.Println(headerErr)
	}
	payloadStr, payloadErr := Base64RawStdEncoding(jwt.Payload)
	if payloadErr != nil {
		log.Println(payloadErr)
	}
	//将经过base64加密后的header和payload部分进行加盐sha256加密形成signature部分
	h := hmac.New(sha256.New, []byte(config.Conf.JsonWebToken.SignatureSecretKey))
	h.Write([]byte(headerStr + "." + payloadStr))
	signatureStr := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	//三部分组成JWT
	jwt_token := headerStr + "." + payloadStr + "." + signatureStr

	return jwt_token
}

// JWT解码
func DecodeTheJWT(jwt_token string) (*model.JWT, string) {
	//拆分JWTToken
	first := strings.Index(jwt_token, ".")
	second := strings.LastIndex(jwt_token, ".")

	headerStr := jwt_token[:first]
	payloadStr := jwt_token[first+1 : second]
	signatureStr := jwt_token[second+1:]

	//对第一部分进行base64解码
	headerData, headerErr := Base64RawStdDecoding(headerStr)
	if headerErr != nil {
		log.Println("jwt heander 解码错误", headerErr)
	}
	var header *model.Header
	json.Unmarshal(headerData, &header)

	//对第二部分(载体)进行base64解码看token是否过期
	payloadData, payloadErr := Base64RawStdDecoding(payloadStr)
	if payloadErr != nil {
		log.Println("jwt payload 解码错误", headerErr)
	}
	var payload *model.Payload
	json.Unmarshal(payloadData, &payload)

	//返回jwt对象和第三部分字符串
	jwt := model.NewJWT(header, payload)

	return jwt, signatureStr
}

// JWT校验
func VerifyTheJWT(jwt *model.JWT, signatureStr string) (bool, string) {
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
	jwtToken := EncodeTheJWT(jwt)
	//拆分JWTToken
	second := strings.LastIndex(jwtToken, ".")

	signatureStrTemp := jwtToken[second+1:]

	return hmac.Equal([]byte(signatureStrTemp), []byte(signatureStr)), "JWT Token校验成功！"
}
