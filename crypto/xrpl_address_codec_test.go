package crypto

import "testing"

func GetUInt32Pointer(value uint32) *uint32 {
	return &value
}

func TestClassicAddressToXAddress(t *testing.T) {
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
		result := ClassicAddressToXAddress(table.rAddress, table.tag, table.test)
		if result != table.xAddress {
			t.Errorf("ClassicAddressToXAddress was incorrect, got: %s, want: %s.", result, table.xAddress)
		}
	}
}

func TestXAddressToClassicAddress(t *testing.T) {
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
		result, tag, test := XAddressToClassicAddress(table.xAddress)
		if result != table.rAddress {
			t.Errorf("XAddressToClassicAddress was incorrect, got: %s, want: %s.", result, table.rAddress)
		}

		if tag == nil || table.tag == nil {
			if tag != table.tag {
				t.Errorf("XAddressToClassicAddress was incorrect, got: %d, want: %d.", tag, table.tag)
			}
		} else if *tag != *(table.tag) {
			t.Errorf("XAddressToClassicAddress was incorrect, got: %d, want: %d.", *tag, *table.tag)
		}

		if test != table.test {
			t.Errorf("XAddressToClassicAddress was incorrect, got: %t, want: %t.", test, table.test)
		}
	}
}
