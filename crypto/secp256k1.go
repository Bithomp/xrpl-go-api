package crypto

import (
	"encoding/binary"
	"fmt"
	"math/big"

	secp "github.com/decred/dcrd/dcrec/secp256k1/v4"
)

const (
	SECP256K1_SEED_LENGTH        = 16
	SECP256K1_PRIVATE_KEY_PREFIX = 0x00
)

var (
	SECP256K1_SEED_PREFIX = []byte{0x21}
)

var (
	order = secp.S256().N
	zero  = big.NewInt(0)
	one   = big.NewInt(1)
)

func GenerateSeedSecp256k1(entropy []byte) (string, error) {
	if entropy == nil {
		var err error
		entropy, err = GenerateSeedEntropy()
		if err != nil {
			return "", err
		}
	}

	encoded := append(SECP256K1_SEED_PREFIX, entropy...)

	return Base58Encode(encoded, ALPHABET), nil
}

func DeriveKeypairSecp256k1(entropy []byte, sequence *uint32) ([]byte, []byte, error) {
	if entropy == nil {
		return nil, nil, fmt.Errorf("entropy is missed")
	}

	if len(entropy) != SECP256K1_SEED_LENGTH {
		return nil, nil, fmt.Errorf("entropy length is not 16")
	}

	// This private generator represents the `root` private key, and is what's
	// used by validators for signing when a keypair is generated from a seed.
	privateGen := deriveScalar(entropy)
	if sequence == nil {
		privateKey := append([]byte{SECP256K1_PRIVATE_KEY_PREFIX}, privateGen.ToECDSA().D.Bytes()...)
		publicKey := privateGen.PubKey().SerializeCompressed()

		return publicKey, privateKey, nil
	}

	// A seed can generate many keypairs as a function of the seed and a uint32.
	// Almost everyone just uses the first account, `0`.
	secpKey := generateKeyForSequence(privateGen.PubKey().SerializeCompressed(), *sequence)
	secpKey.Key = *secpKey.Key.Add(&privateGen.Key)

	privateKey := append([]byte{SECP256K1_PRIVATE_KEY_PREFIX}, secpKey.ToECDSA().D.Bytes()...)
	publicKey := secpKey.PubKey().SerializeCompressed()

	return publicKey, privateKey, nil
}

func deriveScalar(entropy []byte) *secp.PrivateKey {
	inc := big.NewInt(0).SetBytes(entropy)
	inc.Lsh(inc, 32)
	for key := big.NewInt(0); ; inc.Add(inc, one) {
		key.SetBytes(Sha512Half(inc.Bytes()))
		if key.Cmp(zero) > 0 && key.Cmp(order) < 0 {
			return secp.PrivKeyFromBytes(key.Bytes())
		}
	}
}

func generateKeyForSequence(pubKey []byte, sequence uint32) *secp.PrivateKey {
	entropy := make([]byte, secp.PubKeyBytesLenCompressed+4)
	copy(entropy, pubKey)
	binary.BigEndian.PutUint32(entropy[secp.PubKeyBytesLenCompressed:], sequence)

	return deriveScalar(entropy)
}
