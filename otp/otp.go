package otpauth

import (
	"crypto"

	"github.com/kPherox/otpauth/migration"
)

type OTP struct {
	Secret      []byte
	AccountName string
	Issuer      string
	Algorithm   Algorithm
	Digits      int32
}

type HOTP struct {
	OTP
	Counter int64
}

type TOTP struct {
	OTP
	Period int64
}

type OTPAuth interface {
	GetType() OTPType
}

func (_ *HOTP) GetType() OTPType {
	return hotp
}

func (_ *TOTP) GetType() OTPType {
	return totp
}

func MigrationOtpParamToOTPAuth(p *migration.Migration_OtpParameter) (OTPAuth, error) {
	a, err := FromMigrationAlgorithm(p.Algorithm)
	if err != nil {
		return nil, err
	}

	otpauth := &OTP{
		Secret:      p.Secret,
		Issuer:      p.Issuer,
		AccountName: p.Name,
		Algorithm:   a,
		Digits:      p.Digits,
	}

	t, err := FromMigrationOtpType(p.Type)
	switch t {
	case hotp:
		return &HOTP{*otpauth, p.Counter}, nil
	case totp:
		return &TOTP{*otpauth, p.Period}, nil
	default:
		return nil, err
	}
}

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
