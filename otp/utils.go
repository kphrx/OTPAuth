package otpauth

import (
	"encoding/base32"
	"fmt"
	"strings"
)

var digitsPower = []int64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}

func ZeroPadding(v int64, d int) string {
	return fmt.Sprintf("%0*d", d, v%digitsPower[d])
}

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
