package main

import (
	"errors"
	"fmt"
)

func main() {
	// int numbers
	var number int64 = -100000000000
	fmt.Println(number)

	var number2 uint32 = 10000
	fmt.Println(number2)

	// alias
	// INT32 = RUNE
	var number3 rune = 12456
	fmt.Println(number3)

	// BYTE = UINT8
	var number4 byte = 123
	fmt.Println(number4)
	// END INT NUMBERS

	// FLOAT NUMBERS

	var numberFloat1 float32 = 123.45
	fmt.Println(numberFloat1)

	var numberFloat2 float64 = 1230000000.45
	fmt.Println(numberFloat2)

	numberFloat3 := 12345.67
	fmt.Println(numberFloat3)

	// END FLOAT NUMBERS

	// STRINGS
	var str string = "Text"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// END STRINGS

	texto := 5
	fmt.Println(texto)

	var bool1 bool
	fmt.Println(bool1)

	var error error = errors.New("New error")
	fmt.Println(error)
}
