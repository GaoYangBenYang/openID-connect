package utils

import (
	"net/http"
)

// 验证Http Authorization Bearer
func BearerAuth(r *http.Request) (access_token string, ok bool) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", false
	}
	const prefix = "Bearer "
	if len(auth) < len(prefix) {
		return "", false
	}
	c, err := Base64RawURLDecoding(auth[len(prefix):])
	if err != nil {
		return "", false
	}
	access_token = string(c)
	return access_token, true
}
