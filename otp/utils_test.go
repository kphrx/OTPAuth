package otpauth

import (
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

func TestZeroPadding(t *testing.T) {
	run := func(t *testing.T, fixt Fixture, d int) {
		for _, tc := range fixt.TestCases {
			tc := tc
			val, _ := strconv.ParseInt(tc.Value, 10, 64)
			t.Run(fmt.Sprintf("%d to %s", val, tc.Result), func(t *testing.T) {
				t.Parallel()
				if got := ZeroPadding(val, d); got != tc.Result {
					t.Errorf("ZeroPadding(%d, %d) = %s; want %s", val, d, got, tc.Result)
				}
			})
		}
	}
	t.Run("Digits6", func(t *testing.T) {
		t.Parallel()
		run(t, zeroPaddingDigits6, 6)
	})
	t.Run("Digits8", func(t *testing.T) {
		t.Parallel()
		run(t, zeroPaddingDigits8, 8)
	})
}

func TestItob(t *testing.T) {
	for i, tc := range intToBytes.TestCases {
		i, tc := i, tc
		val, _ := strconv.ParseInt(tc.Value, 10, 64)
		t.Run(fmt.Sprintf("%d to %s", val, tc.Result), func(t *testing.T) {
			t.Parallel()
			if b := Itob(int64(val)); fmt.Sprintf("%x", b) != tc.Result || b[7-i] != byte(0x7f) {
				t.Errorf("Itob(%d) = %x; want %s", val, b, tc.Result)
			}
		})
	}
}

func TestBase32Secret(t *testing.T) {
	run := func(t *testing.T, fixt Fixture, f func(*testing.T, string, string)) {
		for _, tc := range fixt.TestCases {
			tc := tc
			t.Run(fmt.Sprintf("%s to %s", tc.Value, tc.Result), func(t *testing.T) {
				t.Parallel()
				f(t, tc.Value, tc.Result)
			})
		}
	}

	t.Run("TestDecodeSecret", func(t *testing.T) {
		t.Parallel()
		decodeSecret := func(t *testing.T, v string, r string) {
			if got, err := DecodeSecret(v); *(*string)(unsafe.Pointer(&got)) != r {
				t.Errorf("DecodeSecret(\"%s\") = %s; want %s", v, got, r)
			} else if err != nil {
				t.Errorf("DecodeSecret(\"%s\") return error: %s", v, err)
			}
		}
		t.Run("NoPadding", func(t *testing.T) {
			t.Parallel()
			run(t, decodeBase32, decodeSecret)
		})
		t.Run("WithPadding", func(t *testing.T) {
			t.Parallel()
			run(t, decodeBase32WithPadding, decodeSecret)
		})
	})

	t.Run("TestEncodeSecret", func(t *testing.T) {
		t.Parallel()
		run(t, encodeBase32, func(t *testing.T, v string, r string) {
			if got := EncodeSecret([]byte(v)); got != r {
				t.Errorf("EncodeSecret(\"%s\") = %s; want %s", v, got, r)
			}
		})
	})
}
