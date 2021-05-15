package main

import (
	"crypto"
	"fmt"
	"net/url"

	"github.com/kpherox/otpauth/otp"
)

func main() {
	u, _ := url.Parse("otpauth://hotp/issuer:OTPAuth?issuer=issuer&secret=GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ&algorithm=SHA1&digits=6&counter=0")
	k, _ := otpauth.DecodeSecret(u.Query().Get("secret"))
	b, _ := otpauth.GenOTP(crypto.SHA1, k, 0)

	fmt.Println(otpauth.ZeroPadding(b, 6))
}
