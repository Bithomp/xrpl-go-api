package crypto

import (
	"fmt"
	"testing"
)

func TestGenerateSeed(t *testing.T) {
	ed25519, _ := GenerateSeed("ed25519")
	if len(ed25519) != 31 {
		t.Errorf("GenerateSeed was incorrect, got: %d, want: %s.", len(ed25519), "31 characters")
	}

	secp256k1, _ := GenerateSeed("secp256k1")
	if len(secp256k1) != 29 {
		t.Errorf("GenerateSeed was incorrect, got: %d, want: %s.", len(secp256k1), "29 characters")
	}
}

func TestDeriveKeypair(t *testing.T) {
	table := []struct {
		seed       string
		publicKey  string
		privateKey string
		rAddress   string
	}{
		{
			"sEdTsfzbUbYKoedPiz7RznJGzBWuGhf",
			"EDF26E95A408796575E4E6B96E8FA33044EDAACE8DC66FC18AEB6D6E1BA0669187",
			"EDDEBDD0E2664B5E340E883F794F965F368C29A0F15668D7AA2F57E519AE1620DF",
			"rJgWgktqTCaGzpPW7JXi5y22k9bdGj8V9d",
		},
		{
			"ssmiqXZXt4GqT8gDWAaE9amftMoUi",
			"03D69B54785F12371F75D5B4DB182571F0C475098DCDA62B543E599CA5EBFF982C",
			"00EED8809A0262B908B4E6E7BC20A1C7ECD21E91646D26D36F5B884A87C7955A19",
			"rAuxnpY9WJdEDsC8f4TEK9CWKvxibD8o6",
		},
	}

	for _, table := range table {
		publicKey, privateKey, _ := DeriveKeypair(table.seed)
		if fmt.Sprintf("%X", publicKey) != table.publicKey {
			t.Errorf("DeriveKeypair was incorrect, got: %s, want: %s.", publicKey, table.publicKey)
		}
		if fmt.Sprintf("%X", privateKey) != table.privateKey {
			t.Errorf("DeriveKeypair was incorrect, got: %s, want: %s.", privateKey, table.privateKey)
		}
	}
}
