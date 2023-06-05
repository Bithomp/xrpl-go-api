package wallet

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	seed, rAddress, _ := Generate("ed25519")

	if len(seed) != 31 {
		t.Errorf("Generate was incorrect, got: %d, want: %s.", len(seed), "31 characters")
	}

	if len(rAddress) != 34 && len(rAddress) != 33 {
		t.Errorf("Generate was incorrect, got: %d, want: %s.", len(rAddress), "33-34 characters")
	}

	seed, rAddress, _ = Generate("secp256k1")

	if len(seed) != 29 {
		t.Errorf("Generate was incorrect, got: %d, want: %s.", len(seed), "29 characters")
	}

	if len(rAddress) != 34 {
		t.Errorf("Generate was incorrect, got: %d, want: %s.", len(rAddress), "34 characters")
	}
}
