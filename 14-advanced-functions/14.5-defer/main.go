package main

import "fmt"

func isStudentApproved(n1, n2 float32) bool {
	defer fmt.Println("Result will be displayed")
	defer fmt.Println("Test")
	fmt.Println("Calculating...")

	avg := (n1 + n2) / 2

	if avg >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isStudentApproved(4, 8))
}
