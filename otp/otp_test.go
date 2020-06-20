package otpauth

import (
	"crypto"
	"fmt"
	"testing"
)

func TestGenOTP(t *testing.T) {
	for i, d := range []int64{
		1284755224,
		1094287082,
		137359152,
		1726969429,
		1640338314,
		868254676,
		1918287922,
		82162583,
		673399871,
		645520489,
	} {
		i, d := i, d
		t.Run(fmt.Sprintf("%d to %d", i, d), func(t *testing.T) {
			t.Parallel()
			if p, _ := GenOTP(crypto.SHA1, []byte("12345678901234567890"), int64(i)); p != d {
				t.Errorf("GenOTP(crypto.SHA1, []byte(\"12345678901234567890\"), %d) = %d; want %d", i, p, d)
			}
		})
	}
}
