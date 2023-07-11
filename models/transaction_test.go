package models

import (
	"testing"
)

func TestEncodeCTID(t *testing.T) {

	table := []struct {
		ledgerIndex uint32
		txIndex     uint16
		networkID   uint16
		ctid        string
	}{
		{1, 2, 3, "C000000100020003"},
		{0, 0, 0, "C000000000000000"},
		{0xFFFFFFF, 0xFFFF, 0xFFFF, "CFFFFFFFFFFFFFFF"},
	}

	for _, table := range table {
		result, _ := EncodeCTID(table.ledgerIndex, table.txIndex, table.networkID)
		if result != table.ctid {
			t.Errorf("EncodeCTID was incorrect, got: %s, want: %s.", result, table.ctid)
		}
	}
}

func TestDecodeCTID(t *testing.T) {
	table := []struct {
		ledgerIndex uint32
		txIndex     uint16
		networkID   uint16
		ctid        string
	}{
		{1, 2, 3, "C000000100020003"},
		{0, 0, 0, "C000000000000000"},
		{0xFFFFFFF, 0xFFFF, 0xFFFF, "CFFFFFFFFFFFFFFF"},
	}

	for _, table := range table {
		ledgerIndex, txIndex, networkID, _ := DecodeCTID(table.ctid)
		if ledgerIndex != table.ledgerIndex {
			t.Errorf("DecodeCTID was incorrect, got: %d, want: %d.", ledgerIndex, table.ledgerIndex)
		}
		if txIndex != table.txIndex {
			t.Errorf("DecodeCTID was incorrect, got: %d, want: %d.", txIndex, table.txIndex)
		}
		if networkID != table.networkID {
			t.Errorf("DecodeCTID was incorrect, got: %d, want: %d.", networkID, table.networkID)
		}
	}

	_, _, _, err := DecodeCTID("C003FFFFFFFFFFFG")
	if err == nil {
		t.Errorf("DecodeCTID should return error")
	}

	_, _, _, err = DecodeCTID("C003FFFFFFFFFFF")
	if err == nil {
		t.Errorf("DecodeCTID should return error")
	}

	_, _, _, err = DecodeCTID("FFFFFFFFFFFFFFFF")
	if err == nil {
		t.Errorf("DecodeCTID should return error")
	}
}
