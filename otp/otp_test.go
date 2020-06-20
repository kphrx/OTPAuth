package otpauth

import (
	"crypto"
	"fmt"
	"testing"
)

func TestZeroPadding(t *testing.T) {
	for n, d := range map[int64]struct {
		Digits  int
		Padding string
	}{
		1234567:   {6, "234567"},
		123456:    {6, "123456"},
		12345:     {6, "012345"},
		987654321: {8, "87654321"},
		98765432:  {8, "98765432"},
		9876543:   {8, "09876543"},
	} {
		n, d := n, d
		t.Run(fmt.Sprintf("Len%d/%d to %s", d.Digits, n, d.Padding), func(t *testing.T) {
			t.Parallel()
			if got := ZeroPadding(n, d.Digits); got != d.Padding {
				t.Errorf("ZeroPadding(%d, %d) = %s; want %s", n, d.Digits, got, d.Padding)
			}
		})
	}
}

func TestGenOTP(t *testing.T) {
	t.Run("UnavailableHashAlgorithm", func(t *testing.T) {
		t.Parallel()
		_, err := GenOTP(crypto.MD5, []byte("12345678901234567890"), 0)
		if err == nil || err.Error() != "Unsupported hash algorithm" {
			t.Fail()
			return
		}
	})

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
