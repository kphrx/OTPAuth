package otpauth

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestItob(t *testing.T) {
	for i, h := range map[int64]string{
		0x7f:               "000000000000007f",
		0x7fff:             "0000000000007fff",
		0x7fffff:           "00000000007fffff",
		0x7fffffff:         "000000007fffffff",
		0x7fffffffff:       "0000007fffffffff",
		0x7fffffffffff:     "00007fffffffffff",
		0x7fffffffffffff:   "007fffffffffffff",
		0x7fffffffffffffff: "7fffffffffffffff",
	} {
		i, h := i, h
		t.Run(fmt.Sprintf("%d to %s", i, h), func(t *testing.T) {
			t.Parallel()
			if b := Itob(i); fmt.Sprintf("%x", b) != h {
				t.Errorf("Itob(%d) = %x; want %s", i, b, h)
			}
		})
	}
}

func TestBase32Secret(t *testing.T) {
	for s, b := range map[string]struct {
		Padding   string
		NoPadding string
	}{
		"f":      {"MY======", "MY"},
		"fo":     {"MZXQ====", "MZXQ"},
		"foo":    {"MZXW6===", "MZXW6"},
		"foob":   {"MZXW6YQ=", "MZXW6YQ"},
		"fooba":  {"MZXW6YTB", "MZXW6YTB"},
		"foobar": {"MZXW6YTBOI======", "MZXW6YTBOI"},
	} {
		s, b := s, b
		t.Run(fmt.Sprintf("TestDecodeSecret/%s", s), func(t *testing.T) {
			t.Parallel()
			t.Run(fmt.Sprintf("WithPadding/%s", b.Padding), func(t *testing.T) {
				t.Parallel()
				if got, err := DecodeSecret(b.Padding); *(*string)(unsafe.Pointer(&got)) != s {
					t.Errorf("DecodeSecret(\"%s\") = %s; want %s", b.Padding, got, s)
				} else if err != nil {
					t.Errorf("DecodeSecret(\"%s\") return error: %s", b.Padding, err)
				}
			})
			t.Run(fmt.Sprintf("NoPadding/%s", b.NoPadding), func(t *testing.T) {
				t.Parallel()
				if got, err := DecodeSecret(b.NoPadding); *(*string)(unsafe.Pointer(&got)) != s {
					t.Errorf("DecodeSecret(\"%s\") = %s; want %s", b.NoPadding, got, s)
				} else if err != nil {
					t.Errorf("DecodeSecret(\"%s\") return error: %s", b.NoPadding, err)
				}
			})
		})
		t.Run(fmt.Sprintf("TestEncodeSecret/%s", s), func(t *testing.T) {
			t.Parallel()
			if got := EncodeSecret([]byte(s)); got != b.NoPadding {
				t.Errorf("EncodeSecret(\"%s\") = %s; want %s", s, got, b.NoPadding)
			}
		})
	}
}
