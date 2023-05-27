package crypto

import (
	"crypto/rand"
	"fmt"
)

const (
	SEED_ENTROPY_LENGTH = 16
	ALGORITHM_ED25519   = "ed25519"
	ALGORITHM_SECP256K1 = "secp256k1"
)

func GenerateSeedEntropy() ([]byte, error) {
	entropy := make([]byte, SEED_ENTROPY_LENGTH)
	_, err := rand.Read(entropy)
	if err != nil {
		return nil, err
	}

	return entropy, nil
}

func GenerateSeed(algorithm string) (string, error) {
	if algorithm == ALGORITHM_ED25519 {
		return GenerateSeedEd25519(nil)
	} else if algorithm == ALGORITHM_SECP256K1 {
		return GenerateSeedSecp256k1(nil)
	} else {
		return "", fmt.Errorf("invalid algorithm")
	}
}

func DecodeSeed(seed string) (string, []byte, error) {
	if seed == "" {
		return "", nil, fmt.Errorf("seed is empty")
	}

	decoded, err := Base58Decode(seed, ALPHABET)
	if err != nil {
		return "", nil, err
	}

	algorithm := ""
	var entropy []byte
	if len(decoded) == 19 && decoded[0] == 0x01 && decoded[1] == 0xe1 && decoded[2] == 0x4b {
		algorithm = ALGORITHM_ED25519
		entropy = decoded[3:]
	} else if len(decoded) == 17 && decoded[0] == 0x21 {
		algorithm = ALGORITHM_SECP256K1
		entropy = decoded[1:]
	} else {
		return "", nil, fmt.Errorf("invalid seed format")
	}

	return algorithm, entropy, nil
}

func DeriveKeypair(seed string) ([]byte, []byte, error) {
	algorithm, entropy, err := DecodeSeed(seed)
	if err != nil {
		return nil, nil, err
	}

	if algorithm == ALGORITHM_ED25519 {
		return DeriveKeypairEd25519(entropy)
	} else if algorithm == ALGORITHM_SECP256K1 {
		sequence := uint32(0)
		return DeriveKeypairSecp256k1(entropy, &sequence)
	} else {
		return nil, nil, fmt.Errorf("invalid algorithm")
	}
}

func GenerateKeyPair(algorithm string) (string, []byte, []byte, error) {
	entropy, err := GenerateSeedEntropy()
	if err != nil {
		return "", nil, nil, err
	}

	var seed string
	var publicKey []byte
	var privateKey []byte
	if algorithm == ALGORITHM_ED25519 {
		seed, err = GenerateSeedEd25519(entropy)
		if err != nil {
			return "", nil, nil, err
		}

		publicKey, privateKey, err = DeriveKeypairEd25519(entropy)
		if err != nil {
			return "", nil, nil, err
		}

	} else if algorithm == ALGORITHM_SECP256K1 {
		seed, err = GenerateSeedSecp256k1(entropy)
		if err != nil {
			return "", nil, nil, err
		}

		sequence := uint32(0)
		publicKey, privateKey, err = DeriveKeypairSecp256k1(entropy, &sequence)
		if err != nil {
			return "", nil, nil, err
		}
	} else {
		return "", nil, nil, fmt.Errorf("invalid algorithm")
	}

	return seed, publicKey, privateKey, nil
}
