package utils

import (
	"encoding/base64"
	"encoding/json"
)

/*
base64编码结果中会有+、/、=三个特殊字符，它们在url中属于特殊字符是直接无法传递的；
base64url其实就是把字符中的'+'和'/'分别替换成'-'和'_'，另外把末尾填充的‘=’去掉;其他都一样。
*/

// base64填充标准编码
func Base64StdEncoding(src interface{}) (string, error) {
	var srcTemp []byte
	switch srcType := src.(type) {
	case []byte:
		srcTemp = srcType
	case string:
		srcTemp = []byte(srcType)
	default:
		var err error
		srcTemp, err = json.Marshal(srcType)
		if err != nil {
			return "", err
		}
	}
	return base64.StdEncoding.EncodeToString(srcTemp), nil
}

// base64填充RUL编码
func Base64URLEncoding(src interface{}) (string, error) {
	var srcTemp []byte
	switch srcType := src.(type) {
	case []byte:
		srcTemp = srcType
	case string:
		srcTemp = []byte(srcType)
	default:
		var err error
		srcTemp, err = json.Marshal(srcType)
		if err != nil {
			return "", err
		}
	}
	return base64.URLEncoding.EncodeToString(srcTemp), nil
}

// base64无填充标准编码
func Base64RawStdEncoding(src interface{}) (string, error) {
	var srcTemp []byte
	switch srcType := src.(type) {
	case []byte:
		srcTemp = srcType
	case string:
		srcTemp = []byte(srcType)
	default:
		var err error
		srcTemp, err = json.Marshal(srcType)
		if err != nil {
			return "", err
		}
	}
	return base64.RawStdEncoding.EncodeToString(srcTemp), nil
}

// base64无填充RUL编码
func Base64RawURLEncoding(src interface{}) (string, error) {
	var srcTemp []byte
	switch srcType := src.(type) {
	case []byte:
		srcTemp = srcType
	case string:
		srcTemp = []byte(srcType)
	default:
		var err error
		srcTemp, err = json.Marshal(srcType)
		if err != nil {
			return "", err
		}
	}
	return base64.RawURLEncoding.EncodeToString(srcTemp), nil
}

// base64填充标准解码
func Base64StdDecoding(src string) ([]byte, error) {
	srcTemp, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}
	return srcTemp, nil
}

// base64填充RUL解码
func Base64URLDecoding(src string) ([]byte, error) {
	srcTemp, err := base64.URLEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}
	return srcTemp, nil
}

// base64无填充标准解码
func Base64RawStdDecoding(src string) ([]byte, error) {
	srcTemp, err := base64.RawStdEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}
	return srcTemp, nil
}

// base64无填充RUL解码
func Base64RawURLDecoding(src string) ([]byte, error) {
	srcTemp, err := base64.RawURLEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}
	return srcTemp, nil
}
