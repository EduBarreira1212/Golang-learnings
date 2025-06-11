package main

import "fmt"

func Sum(num1, num2 int8) (int8, int8) {
	sum := num1 + num2
	sub := num1 - num2

	return sum, sub
}

var f = func(text string) string {
	fmt.Println(text)
	return text
}

func main() {
	sumResult, subResult := Sum(10, 5)
	fmt.Println(sumResult, subResult)

	_, subResult2 := Sum(8, 4)
	fmt.Println(subResult2)

	f("Test")
}
