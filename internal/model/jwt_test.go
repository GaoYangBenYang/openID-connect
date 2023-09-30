package model

import (
	"fmt"
	"testing"
)

func TestEncodeTheJWT(t *testing.T) {

}

func TestDecodeTheJWT(t *testing.T) {
	
}

func TestVerifyTheJWT(t *testing.T) {
	jwta, stra := DecodeTheJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJsb2NhbGhvc3Q6ODAwMCIsInN1YiI6Ik9wZW5JRFByb3ZpZGVyIiwiYXVkIjoibG9jYWxob3N0OjgwODEiLCJuYmYiOjE2OTYwNDE1NDQsImV4cCI6MTY5NjA0MTg0NCwiaWF0IjoxNjk2MDQxNTQ0LCJqdGkiOiJKV1QgSUQgMSIsInNlc3Npb25fc3RhdGUiOiLpq5jmtIsifQ.XPEsbQghEyso9OdKwSXm4-7TqVIlB-27O-9bN2jQ-0c")
	fmt.Println(VerifyTheJWT(jwta, stra))
}
