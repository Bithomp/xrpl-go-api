package address_codec

import (
	"testing"
)

func getUInt32Pointer(value uint32) *uint32 {
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
		{"rHiJahydBswnAUMZk5yhTjTvcjBE1fXAGh", nil, false, "XVQT4qc3xZCA2agKHNTvqRNMKknq8BDqhRnEX6o9mV1GPC5"},
		{"rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", getUInt32Pointer(123), false, "X721hAAhjjN1AksX7aZLLtFG6Ks4zy63VEmzcFBdNmTdWUa"},
		{"rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", getUInt32Pointer(0), false, "X721hAAhjjN1AksX7aZLLtFG6Ks4zyeo1pUFGC3HrjJL2p4"},
		{"rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", getUInt32Pointer(123), true, "T7WPVhasbg8uPEF3juJfNt2Bwk2BWZMwSRjSvVR9qMW7kdu"},
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
		{"X721hAAhjjN1AksX7aZLLtFG6Ks4zy63VEmzcFBdNmTdWUa", "rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", getUInt32Pointer(123), false},
		{"X721hAAhjjN1AksX7aZLLtFG6Ks4zyeo1pUFGC3HrjJL2p4", "rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", getUInt32Pointer(0), false},
		{"T7WPVhasbg8uPEF3juJfNt2Bwk2BWZMwSRjSvVR9qMW7kdu", "rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", getUInt32Pointer(123), true},
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

func TestNodePublicToClassicAddress(t *testing.T) {
	table := []struct {
		nodePublic string
		rAddress   string
	}{
		{"nHBtDzdRDykxiuv7uSMPTcGexNm879RUUz5GW4h1qgjbtyvWZ1LE", "rHiJahydBswnAUMZk5yhTjTvcjBE1fXAGh"},
	}

	for _, table := range table {
		result, err := NodePublicToClassicAddress(table.nodePublic)
		if err != nil {
			t.Errorf("NodePublicToClassicAddress was incorrect, got: %v, want: %v.", err, nil)
		}

		if result != table.rAddress {
			t.Errorf("NodePublicToClassicAddress was incorrect, got: %s, want: %s.", result, table.rAddress)
		}
	}
}

func TestNodePublicToXAddress(t *testing.T) {
	table := []struct {
		nodePublic string
		xAddress   string
		tag        *uint32
		test       bool
	}{
		{"nHBtDzdRDykxiuv7uSMPTcGexNm879RUUz5GW4h1qgjbtyvWZ1LE", "XVQT4qc3xZCA2agKHNTvqRNMKknq8BDqhRnEX6o9mV1GPC5", nil, false},
		{"nHBtDzdRDykxiuv7uSMPTcGexNm879RUUz5GW4h1qgjbtyvWZ1LE", "TVKy8AmfFXDyrcSvZsU6jQzEoBceHk11rbp6qPpg4sxWHCg", nil, true},
	}

	for _, table := range table {
		result := NodePublicToXAddress(table.nodePublic, table.tag, table.test)
		if result != table.xAddress {
			t.Errorf("NodePublicToXAddress was incorrect, got: %s, want: %s.", result, table.xAddress)
		}
	}
}

func TestDecodeAccountID(t *testing.T) {
	table := []struct {
		rAddress  string
		accountId []byte
	}{
		{"rsuUjfWxrACCAwGQDsNeZUhpzXf1n1NK5Z", []byte{0x1f, 0xcf, 0xab, 0x86, 0x1b, 0xb2, 0x98, 0x3f, 0x7c, 0x7d, 0xca, 0x75, 0x1b, 0xff, 0xd2, 0x89, 0x76, 0x52, 0x89, 0x09}},
		{"rsEUPfWoWZNikoqr9ZZg7vPpNgPG9XXB43", []byte{0x18, 0x9f, 0x93, 0x5b, 0xc2, 0x61, 0xb3, 0xbc, 0x83, 0x25, 0x3b, 0xe0, 0x4a, 0x2f, 0xca, 0x81, 0x1b, 0xa2, 0xc4, 0x80}},
	}

	for _, table := range table {
		result, err := DecodeAccountID(table.rAddress)
		if err != nil {
			t.Errorf("DecodeAccountID was incorrect, got: %v, want: %v.", err, nil)
		}

		if len(result) != len(table.accountId) {
			t.Errorf("DecodeAccountID was incorrect, got: %d, want: %d.", len(result), len(table.accountId))
		}

		for i := 0; i < len(result); i++ {
			if result[i] != table.accountId[i] {
				t.Errorf("DecodeAccountID was incorrect for byte %d, got: %d, want: %d.", i, result[i], table.accountId[i])
				break
			}
		}
	}
}
