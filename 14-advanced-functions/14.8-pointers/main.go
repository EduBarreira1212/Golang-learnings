package main

import "fmt"

func invertNumber(number int) {
	number = number * -1
}

func invertNumberWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 10
	invertNumber(number)
	fmt.Println(number)

	fmt.Println("-------------")

	number2 := 50
	fmt.Println(number2)
	invertNumberWithPointer(&number2)
	fmt.Println(number2)
}
