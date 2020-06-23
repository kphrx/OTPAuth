package otpauth

import (
	"crypto"
	"fmt"
	"strconv"
	"testing"
)

func TestGenOTP(t *testing.T) {
	t.Run("UnavailableHashAlgorithm", func(t *testing.T) {
		t.Parallel()
		_, err := GenOTP(crypto.MD5, []byte("12345678901234567890"), 0)
		if err == nil || err.Error() != "Unsupported hash algorithm" {
			t.Fail()
			return
		}
	})

	for i, tc := range hotpValue.TestCases {
		i := i
		val, _ := strconv.ParseInt(tc.Result, 10, 64)
		t.Run(fmt.Sprintf("%d to %d", i, val), func(t *testing.T) {
			t.Parallel()
			if p, _ := GenOTP(crypto.SHA1, []byte("12345678901234567890"), int64(i)); p != val {
				t.Errorf("GenOTP(crypto.SHA1, []byte(\"12345678901234567890\"), %d) = %d; want %d", i, p, val)
			}
		})
	}
}
