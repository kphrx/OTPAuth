package otpauth

import (
	"crypto"
	"fmt"
	"testing"
)

func TestHMAC(t *testing.T) {
	key := []byte("12345678901234567890")

	t.Run("UnavailableHashAlgorithm", func(t *testing.T) {
		t.Parallel()
		_, err := HMAC(crypto.MD5, key, Itob(0))
		if err == nil || err.Error() != "Unsupported hash algorithm" {
			t.Fail()
			return
		}
	})

	for i, tc := range hmacHash.TestCases {
		i, tc := i, tc
		t.Run(fmt.Sprintf("%d to %s", i, tc.Result), func(t *testing.T) {
			t.Parallel()
			if b, _ := HMAC(crypto.SHA1, key, Itob(int64(i))); fmt.Sprintf("%x", b) != tc.Result {
				t.Errorf("HMAC(crypto.SHA1, []byte(\"12345678901234567890\"), Itob(%d)) = %x; want %s", i, b, tc.Result)
			}
		})
	}
}
