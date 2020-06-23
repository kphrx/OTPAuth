package otpauth

import (
	"crypto"
)

func GenOTP(a crypto.Hash, s []byte, f int64) (int64, error) {
	h, err := HMAC(a, s, Itob(f))
	if err != nil {
		return 0, err
	}

	o := h[len(h)-1] & 0xf
	b := ((int64(h[o]) & 0x7f) << 24) |
		((int64(h[o+1]) & 0xff) << 16) |
		((int64(h[o+2]) & 0xff) << 8) |
		(int64(h[o+3]) & 0xff)

	return b, nil
}
