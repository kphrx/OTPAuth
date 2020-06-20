package main

import (
	"crypto"
	"fmt"
	"net/url"

	"github.com/kPherox/otpauth/otp"
)

func main() {
	u, _ := url.Parse("otpauth://totp/issuer:OTPAuth?issuer=issuer&secret=GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ&algorithm=SHA1&digits=6&period=30")
	k, _ := otpauth.DecodeSecret(u.Query().Get("secret"))
	hs, _ := otpauth.HMAC(crypto.SHA1, k, otpauth.Itob(0))
	o := hs[len(hs)-1] & 0xf
	b := ((int64(hs[o]) & 0x7f) << 24) |
		((int64(hs[o+1]) & 0xff) << 16) |
		((int64(hs[o+2]) & 0xff) << 8) |
		(int64(hs[o+3]) & 0xff)

	fmt.Printf("%06d\n", b%1000000)
}
