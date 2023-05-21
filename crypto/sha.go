package crypto

import (
	"crypto/sha256"
)

// Based on https://github.com/rubblelabs/ripple/blob/master/crypto/util.go

func DoubleSha256(b []byte) []byte {
	hasher := sha256.New()
	hasher.Write(b)
	sha := hasher.Sum(nil)
	hasher.Reset()
	hasher.Write(sha)
	return hasher.Sum(nil)
}
