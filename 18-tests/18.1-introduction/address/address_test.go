package address

import "testing"

type TestingStruct struct {
	addressForTest string
	expectedReturn string
}

func TestVerifyAddressType(t *testing.T) {

	valuesForTest := []TestingStruct{
		{"Street 1", "STREET"},
		{"Avenue 123", "AVENUE"},
		{"Block 10", "Invalid type"},
		{"street 12345", "STREET"},
		{"aVENUE test", "AVENUE"},
		{"AVENUE test12345", "AVENUE"},
	}

	for _, caseTest := range valuesForTest {
		typeReturned := VerifyAddressType(caseTest.addressForTest)

		if typeReturned != caseTest.expectedReturn {
			t.Errorf("Expected: %s; Returned: %s", caseTest.expectedReturn, typeReturned)
		}
	}
}
