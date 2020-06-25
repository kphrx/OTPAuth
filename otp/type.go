package otpauth

import (
	"errors"

	"github.com/kPherox/otpauth/migration"
)

type OTPType int

const (
	hotp = iota + 1
	totp
)

func FromMigrationOtpType(ma migration.Migration_OtpType) (Algorithm, error) {
	switch ma {
	case migration.Migration_OTP_HOTP:
		return hotp, nil
	case migration.Migration_OTP_TOTP:
		return totp, nil
	default:
		return 0, errors.New("Invalid otp algorithm")
	}
}

func (a *OTPType) ToMigrationOtpType() migration.Migration_OtpType {
	switch *a {
	case hotp:
		return migration.Migration_OTP_HOTP
	case totp:
		return migration.Migration_OTP_TOTP
	default:
		return migration.Migration_OTP_INVALID
	}
}
