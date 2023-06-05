package wallet

import (
	"github.com/Bithomp/xrpl-go-api/crypto"
)

func Generate(algorithm string) (string, string, error) {
	seed, publicKey, _, err := crypto.GenerateKeyPair(algorithm)
	if err != nil {
		return "", "", err
	}

	// generate account id
	accountId := crypto.Sha256RipeMD160(publicKey)
	rAddress := crypto.EncodeAccountID(accountId)

	return seed, rAddress, nil
}
