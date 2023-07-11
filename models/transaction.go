package models

import "fmt"

func EncodeCTID(ledgerIndex uint32, txIndex uint16, networkID uint16) (string, error) {
	if ledgerIndex > 0xFFFFFFF {
		return "", fmt.Errorf("ledgerIndex is too large")
	}

	if txIndex > 0xFFFF {
		return "", fmt.Errorf("txIndex is too large")
	}

	if networkID > 0xFFFF {
		return "", fmt.Errorf("networkID is too large")
	}

	var ctidValue uint64 = 0xC000000000000000
	ctidValue = ctidValue + uint64(ledgerIndex)<<32 | uint64(txIndex)<<16 | uint64(networkID)

	hex := fmt.Sprintf("%X", ctidValue)

	return hex, nil
}

func DecodeCTID(ctid string) (ledgerIndex uint32, txIndex uint16, networkID uint16, err error) {
	if len(ctid) != 16 {
		return 0, 0, 0, fmt.Errorf("invalid CTID")
	}

	var ctidValue uint64
	_, err = fmt.Sscanf(ctid, "%X", &ctidValue)
	if err != nil {
		return 0, 0, 0, err
	}

	if ctidValue>>60 != 0xC {
		return 0, 0, 0, fmt.Errorf("invalid CTID")
	}

	ledgerIndex = uint32(ctidValue >> 32 & 0xFFFFFFF)
	txIndex = uint16(ctidValue >> 16 & 0xFFFF)
	networkID = uint16(ctidValue & 0xFFFF)

	return ledgerIndex, txIndex, networkID, nil
}
