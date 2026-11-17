package mptoken

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/Bithomp/xrpl-go-api/address_codec"
)

var (
	ErrInvalidMPTokenIssuanceID = errors.New("invalid MPToken issuance ID")
)

func BuildMPTokenIssuanceID(sequence uint32, issuer string) (string, error) {
	// Convert sequence to 8-char uppercase hex
	seqHex := fmt.Sprintf("%08X", sequence)

	// Decode issuer address to account ID (20 bytes)
	accountID, err := address_codec.DecodeAccountID(issuer)
	if err != nil {
		return "", err
	}

	// Convert accountID to hex and make uppercase
	accountHex := hex.EncodeToString(accountID)
	accountHex = strings.ToUpper(accountHex)

	fmt.Printf("Account Hex: %s\n", accountHex)

	return seqHex + accountHex, nil
}

func ParseMPTokenIssuanceID(mptIssuanceID string) (sequence uint32, issuer string, err error) {
	if len(mptIssuanceID) != 48 {
		return 0, "", ErrInvalidMPTokenIssuanceID
	}

	// Parse sequence
	seqHex := mptIssuanceID[:8]
	seqInt := new(big.Int)
	seqInt.SetString(seqHex, 16)
	sequence = uint32(seqInt.Uint64())

	// Parse issuer
	accountHex := mptIssuanceID[8:]
	accountBytes, err := hex.DecodeString(accountHex)
	if err != nil {
		return 0, "", err
	}
	issuer = address_codec.EncodeAccountID(accountBytes)

	return sequence, issuer, nil
}
