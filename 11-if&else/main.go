package main

import "fmt"

func main() {
	number := 1

	if number > 5 {
		fmt.Println("Number greater than 5")
	} else {
		fmt.Println("Number smaller than 5")
	}

	if otherNumber := number; otherNumber > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}
}
