package otpauth

import (
	"crypto"
	"crypto/hmac"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"errors"
)

func HMAC(a crypto.Hash, k []byte, t []byte) ([]byte, error) {
	if !a.Available() {
		return nil, errors.New("Unsupported hash algorithm")
	}

	hsr := hmac.New(a.New, k)
	hsr.Write(t)
	return hsr.Sum(nil), nil
}
