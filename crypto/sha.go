package crypto

import (
	"crypto/sha256"
	"crypto/sha512"

	//lint:ignore SA1019 we need ripemd160
	"golang.org/x/crypto/ripemd160"
)

// Based on https://github.com/rubblelabs/ripple/blob/master/crypto/util.go

// Returns bytes of a SHA256 of the input bytes
func Sha256(b []byte) []byte {
	hasher := sha256.New()
	hasher.Write(b)
	return hasher.Sum(nil)
}

// Returns bytes of a SHA256 double hash of the input bytes
func DoubleSha256(b []byte) []byte {
	hasher := sha256.New()
	hasher.Write(b)
	sha := hasher.Sum(nil)
	hasher.Reset()
	hasher.Write(sha)
	return hasher.Sum(nil)
}

// Returns bytes of a SHA256 ripemd160 of the input bytes
func Sha256RipeMD160(b []byte) []byte {
	ripe := ripemd160.New()
	sha := sha256.New()
	sha.Write(b)
	ripe.Write(sha.Sum(nil))
	return ripe.Sum(nil)
}

// Returns first 32 bytes of a SHA512 of the input bytes
func Sha512Half(b []byte) []byte {
	hasher := sha512.New()
	hasher.Write(b)
	return hasher.Sum(nil)[:32]
}
