package address

import "testing"

func TestVerifyAddressType(t *testing.T) {
	addressForTesting := "Street 12345"
	typeExpected := "STREET"
	typeReturned := VerifyAddressType(addressForTesting)

	if typeReturned != typeExpected {
		t.Errorf("Expected: %s; Returned: %s", typeExpected, typeReturned)
	}
}
