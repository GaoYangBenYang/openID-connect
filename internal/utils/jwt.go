package utils

import (
	"OpenIDProvider/internal/middleware"
	"OpenIDProvider/internal/model"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// JWT编码
func EncodeTheJWT(jwt *model.JWT) (string, error) {
	//Base64URL算法进行header和payload转化字符串
	headerStr, headerErr := Base64RawURLEncoding(jwt.Header)
	if headerErr != nil {
		return "", errors.New("对JWT Header进行编码时发生错误")
	}
	payloadStr, payloadErr := Base64RawURLEncoding(jwt.Payload)
	if payloadErr != nil {
		return "", errors.New("对JWT Payload进行编码时发生错误")
	}
	//将经过base64加密后的header和payload部分进行加盐sha256加密形成signature部分
	h := hmac.New(sha256.New, []byte(middleware.Config.JsonWebToken.SignatureSecretKey))
	h.Write([]byte(headerStr + "." + payloadStr))
	signatureStr, signatureErr := Base64RawURLEncoding(h.Sum(nil))
	if signatureErr != nil {
		return "", errors.New("对JWT Signature进行编码时发生错误")
	}
	return headerStr + "." + payloadStr + "." + signatureStr, nil
}

// JWT解码
func DecodeTheJWT(json_web_token string) (*model.JWT, error) {
	//拆分json_web_token
	first := strings.Index(json_web_token, ".")
	second := strings.LastIndex(json_web_token, ".")
	headerStr := json_web_token[:first]
	payloadStr := json_web_token[first+1 : second]
	signatureStr := json_web_token[second+1:]

	//对Header部分进行base64无填充标准解码
	headerData, headerErr := Base64RawURLDecoding(headerStr)
	if headerErr != nil {
		return nil, errors.New("对JWT Header进行解码时发生错误")
	}
	var header *model.Header
	json.Unmarshal(headerData, &header)

	//对Payload部分进行base64无填充标准解码
	payloadData, payloadErr := Base64RawURLDecoding(payloadStr)
	if payloadErr != nil {
		return nil, errors.New("对JWT Payload进行解码时发生错误")
	}
	var payload *model.Payload
	json.Unmarshal(payloadData, &payload)

	//获取当前unix时间戳
	timeUnix := int(time.Now().Unix())
	//校验生效时间是否到达
	if timeUnix < payload.Nbf {
		return nil, errors.New("JWT生效时间未到达")
	}
	//校验JWT有效时间是否过期
	if timeUnix > payload.Exp {
		return nil, errors.New("JWT已过期")
	}

	//解析JWT对象
	jwt := model.NewJWT(header, payload)

	//对解析出来的JWT对象进行编码并进行JWT校验
	jwtStr, jwtErr := EncodeTheJWT(jwt)
	if jwtErr != nil {
		return nil, jwtErr
	}
	//拆分JWT的
	signatureStrTemp := jwtStr[strings.LastIndex(jwtStr, ".")+1:]
	//Equal比较加密字符串
	if hmac.Equal([]byte(signatureStrTemp), []byte(signatureStr)) {
		return jwt, nil
	}
	return nil, errors.New("JWT校验不通过")
}
