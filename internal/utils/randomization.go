package utils

import (
	"math/rand"
	"strings"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 随机生成Code
func RandomCode() string {
	const N = 22
	sb := strings.Builder{}
	sb.Grow(N)
	for i := 0; i < N; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}