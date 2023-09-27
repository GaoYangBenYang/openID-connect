package model

import (
	"fmt"
	"testing"
)

func TestEncodeTheJWT(t *testing.T) {
	header := NewHeader("HS256", "JWT")
	payload := NewPayload("www.bejson.com", "evander", "audrey", 1695827196, 1695929385, 1695827196, "", "高洋")
	signature := NewSignature("asdas")
	jwt := NewJWT(header, payload, signature)
	fmt.Println(jwt.EncodeTheJWT())

}
