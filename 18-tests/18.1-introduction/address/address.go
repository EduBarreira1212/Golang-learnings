package address

import "strings"

func VerifyAddressType(address string) string {
	validTypes := []string{"street", "avenue"}

	lowerCaseAddress := strings.ToLower(address)
	firstWordOfAddress := strings.Split(lowerCaseAddress, " ")[0]

	isAddressValid := false

	for _, validType := range validTypes {
		if validType == firstWordOfAddress {
			isAddressValid = true
		}
	}

	if isAddressValid {
		return strings.ToTitle(firstWordOfAddress)
	}

	return "Invalid type"
}
