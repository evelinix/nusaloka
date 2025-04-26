package helper

import (
	"encoding/base64"
)

func Base64URLEncode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func Base64URLDecode(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}