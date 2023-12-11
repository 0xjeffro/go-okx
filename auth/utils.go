package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func HmacSHA256(data, key []byte) string {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
