package main

import (
	"fmt"
	"test/testing"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Packages")
	testing.Write()

	response := checkmail.ValidateFormat("edu@gmail.com")
	fmt.Println(response)
}
