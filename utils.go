package otpauth

import (
	"encoding/base32"
	"strings"
)

func Itob(n int64) []byte {
	b := make([]byte, 8)
	for i := range b {
		b[7-i] = byte(n & 0xff)
		n >>= 8
	}
	return b
}

var base32Encoding = base32.StdEncoding.WithPadding(base32.NoPadding)

func DecodeSecret(s string) ([]byte, error) {
	s = strings.ToUpper(s)
	s = strings.TrimRight(s, "=")
	return base32Encoding.DecodeString(s)
}

func EncodeSecret(s []byte) string {
	return base32Encoding.EncodeToString(s)
}
