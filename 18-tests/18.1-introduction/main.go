package main

import (
	"fmt"
	address "tests-introduction/address"
)

func main() {
	addressType := address.VerifyAddressType("Street 123")
	addressType2 := address.VerifyAddressType("street 1")
	addressType3 := address.VerifyAddressType("Avenue 123")
	addressType4 := address.VerifyAddressType("avenue 00000")

	fmt.Println(addressType)
	fmt.Println(addressType2)
	fmt.Println(addressType3)
	fmt.Println(addressType4)
}
