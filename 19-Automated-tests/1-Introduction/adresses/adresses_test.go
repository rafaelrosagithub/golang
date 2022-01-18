package adresses

import "testing"

// Unit Test
func TestAddressType(t *testing.T) {
	testAddress := "Avenue Paulista"
	expectedAddressType := "Avenue"
	receivedAddressType := AddressType(testAddress)

	if receivedAddressType != expectedAddressType {
		t.Errorf("Received type is different from expected! Expected %s and received %s",
			expectedAddressType,
			receivedAddressType)
	}
}
