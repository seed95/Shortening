package random

import (
	"math/rand"
	"time"
)

var _charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func StringWithCharset(length uint, charset string) string {
	if len(charset) == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func String(length uint) string {
	return StringWithCharset(length, _charset)
}
