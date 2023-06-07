package address_codec

import (
	"fmt"

	"github.com/Bithomp/xrpl-go-api/crypto"
)

const (
	ACCOUNT_ID = 0
)

var (
	// 5, 68
	X_ADDRESS_PREFIX_BYTES_MAIN = []byte{0x05, 0x44}
	// 4, 147
	X_ADDRESS_PREFIX_BYTES_TEST = []byte{0x04, 0x93}
)

// https://github.com/XRPLF/xrpl.js/blob/main/packages/ripple-address-codec/src/index.ts#L25
func ClassicAddressToXAddress(rAddress string, tag *uint32, test bool) string {
	accountId, err := DecodeAccountID(rAddress)
	if err != nil {
		return ""
	}

	return EncodeXAddress(accountId, tag, test)
}

// https://github.com/XRPLF/xrpl.js/blob/main/packages/ripple-address-codec/src/index.ts#L76
func XAddressToClassicAddress(xAddress string) (string, *uint32, bool) {
	decoded, err := crypto.Base58Decode(xAddress, crypto.ALPHABET)
	if err != nil {
		return "", nil, false
	}

	// check length
	if len(decoded) != 31 {
		return "", nil, false
	}

	test := false

	// version and network
	if decoded[0] == X_ADDRESS_PREFIX_BYTES_MAIN[0] && decoded[1] == X_ADDRESS_PREFIX_BYTES_MAIN[1] {
		test = false
	} else if decoded[0] == X_ADDRESS_PREFIX_BYTES_TEST[0] && decoded[1] == X_ADDRESS_PREFIX_BYTES_TEST[1] {
		test = true
	} else {
		return "", nil, false
	}

	// get account id
	accountId := decoded[2:22]

	// get tag flag
	tagFlag := decoded[22]

	// get tag
	theTag := uint32(decoded[23]) | uint32(decoded[24])<<8 | uint32(decoded[25])<<16 | uint32(decoded[26])<<24

	tag := (*uint32)(nil)
	if tagFlag == 1 {
		tag = &theTag
	}

	return EncodeAccountID(accountId), tag, test
}

func EncodeAccountID(accountId []byte) string {
	// add payload length, 1 byte
	encoded := append([]byte{byte(ACCOUNT_ID)}, accountId...)

	return crypto.Base58Encode(encoded, crypto.ALPHABET)
}

func DecodeAccountID(rAddress string) ([]byte, error) {
	decoded, err := crypto.Base58Decode(rAddress, crypto.ALPHABET)
	if err != nil {
		return nil, err
	}

	// remove payload length, 1 byte
	decoded = decoded[1:]

	if len(decoded) != 20 {
		return nil, fmt.Errorf("invalid account id length")
	}

	return decoded, nil
}

func EncodeXAddress(accountId []byte, tag *uint32, test bool) string {
	prefixBytes := X_ADDRESS_PREFIX_BYTES_MAIN
	if test {
		prefixBytes = X_ADDRESS_PREFIX_BYTES_TEST
	}

	tagFlag := byte(0)
	if tag != nil {
		tagFlag = byte(1)
	}

	theTag := uint32(0)
	if tag != nil {
		theTag = *tag
	}

	// add version and network
	encoded := append(prefixBytes, accountId...)
	// add tag flag
	encoded = append(encoded, tagFlag)
	// add tag
	encoded = append(encoded, byte(theTag), byte(theTag>>8), byte(theTag>>16), byte(theTag>>24))
	// add reserved bytes (4) when tag will be uint64
	encoded = append(encoded, byte(0), byte(0), byte(0), byte(0))

	return crypto.Base58Encode(encoded, crypto.ALPHABET)
}
