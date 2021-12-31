package shortening

import (
	"crypto/md5"
	"encoding/base64"
)

func hash(data string) []byte {
	digest := md5.New()
	digest.Write([]byte(data))
	return digest.Sum(nil)
}

func encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func decode(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}
