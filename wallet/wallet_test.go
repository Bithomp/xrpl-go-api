package wallet

import (
	"testing"
)

func GetUInt32Pointer(value uint32) *uint32 {
	return &value
}

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

func TestRAddressToXAddress(t *testing.T) {
	table := []struct {
		rAddress string
		tag      *uint32
		test     bool
		xAddress string
	}{
		{"rsuUjfWxrACCAwGQDsNeZUhpzXf1n1NK5Z", nil, false, "X7czuu79XJ4GHhN5bsHDNyNjCrDFgjXw9rE9ELS86d47DXo"},
		{"rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", GetUInt32Pointer(123), false, "X721hAAhjjN1AksX7aZLLtFG6Ks4zy63VEmzcFBdNmTdWUa"},
		{"rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", GetUInt32Pointer(0), false, "X721hAAhjjN1AksX7aZLLtFG6Ks4zyeo1pUFGC3HrjJL2p4"},
		{"rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", GetUInt32Pointer(123), true, "T7WPVhasbg8uPEF3juJfNt2Bwk2BWZMwSRjSvVR9qMW7kdu"},
	}

	for _, table := range table {
		result := RAddressToXAddress(table.rAddress, table.tag, table.test)
		if result != table.xAddress {
			t.Errorf("RAddressToXAddress was incorrect, got: %s, want: %s.", result, table.xAddress)
		}
	}
}

func TestXAddressToRAddress(t *testing.T) {
	table := []struct {
		xAddress string
		rAddress string
		tag      *uint32
		test     bool
	}{
		{"X7czuu79XJ4GHhN5bsHDNyNjCrDFgjXw9rE9ELS86d47DXo", "rsuUjfWxrACCAwGQDsNeZUhpzXf1n1NK5Z", nil, false},
		{"X721hAAhjjN1AksX7aZLLtFG6Ks4zy63VEmzcFBdNmTdWUa", "rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", GetUInt32Pointer(123), false},
		{"X721hAAhjjN1AksX7aZLLtFG6Ks4zyeo1pUFGC3HrjJL2p4", "rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", GetUInt32Pointer(0), false},
		{"T7WPVhasbg8uPEF3juJfNt2Bwk2BWZMwSRjSvVR9qMW7kdu", "rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", GetUInt32Pointer(123), true},
	}

	for _, table := range table {
		result, tag, test := XAddressToRAddress(table.xAddress)
		if result != table.rAddress {
			t.Errorf("XAddressToRAddress was incorrect, got: %s, want: %s.", result, table.rAddress)
		}

		if tag == nil || table.tag == nil {
			if tag != table.tag {
				t.Errorf("XAddressToRAddress was incorrect, got: %d, want: %d.", tag, table.tag)
			}
		} else if *tag != *(table.tag) {
			t.Errorf("XAddressToRAddress was incorrect, got: %d, want: %d.", *tag, *table.tag)
		}

		if test != table.test {
			t.Errorf("XAddressToRAddress was incorrect, got: %t, want: %t.", test, table.test)
		}
	}
}
