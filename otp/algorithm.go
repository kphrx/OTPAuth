package otpauth

import (
	"crypto"
	"errors"

	"github.com/kPherox/otpauth/migration"
)

type Algorithm int

const (
	sha1 Algorithm = iota + 1
	sha256
	sha512
)

func FromMigrationAlgorithm(ma migration.Migration_Algorithm) (Algorithm, error) {
	switch ma {
	case migration.Migration_ALGO_SHA1:
		return sha1, nil
	case migration.Migration_ALGO_SHA256:
		return sha256, nil
	case migration.Migration_ALGO_SHA512:
		return sha512, nil
	default:
		return 0, errors.New("Unspecified hash algorithm")
	}
}

func (a *Algorithm) ToMigrationAlgorithm() migration.Migration_Algorithm {
	switch *a {
	case sha1:
		return migration.Migration_ALGO_SHA1
	case sha256:
		return migration.Migration_ALGO_SHA256
	case sha512:
		return migration.Migration_ALGO_SHA512
	default:
		return migration.Migration_ALGO_UNSPECIFIED
	}
}

func (a *Algorithm) Hash() crypto.Hash {
	switch *a {
	case sha1:
		return crypto.SHA1
	case sha256:
		return crypto.SHA256
	case sha512:
		return crypto.SHA512
	default:
		return crypto.SHA1
	}
}
