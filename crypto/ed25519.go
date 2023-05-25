package crypto

import (
	"bytes"
	"crypto/ed25519"
	"fmt"
)

const (
	ED25519_SEED_LENGTH        = 16
	ED25519_PRIVATE_KEY_PREFIX = 0xED
)

var (
	ED25519_SEED_PREFIX = []byte{0x01, 0xe1, 0x4b}
)

func GenerateSeedEd25519(entropy []byte) (string, error) {
	if entropy == nil {
		var err error
		entropy, err = GenerateSeedEntropy()
		if err != nil {
			return "", err
		}
	}

	encoded := append(ED25519_SEED_PREFIX, entropy...)

	return Base58Encode(encoded, ALPHABET), nil
}

func DeriveKeypairEd25519(entropy []byte) ([]byte, []byte, error) {
	if entropy == nil {
		return nil, nil, fmt.Errorf("entropy is nil")
	}

	if len(entropy) != ED25519_SEED_LENGTH {
		return nil, nil, fmt.Errorf("entropy length is not 16")
	}

	r := bytes.NewReader(Sha512Half(entropy))

	// publicKey is the same as privateKey[32:]
	publicKey, privateKey, err := ed25519.GenerateKey(r)
	if err != nil {
		return nil, nil, err
	}

	publicKey = append([]byte{ED25519_PRIVATE_KEY_PREFIX}, publicKey...)
	privateKey = append([]byte{ED25519_PRIVATE_KEY_PREFIX}, privateKey[:32]...)

	return publicKey, privateKey, nil
}
