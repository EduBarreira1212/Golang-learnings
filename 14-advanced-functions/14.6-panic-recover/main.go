package main

import "fmt"

func recoverExec() {
	if r := recover(); r != nil {
		fmt.Println("Execution recovered with sucess!")
	}
}

func isStudentApproved(n1, n2 float64) bool {
	defer recoverExec()
	avg := (n1 + n2) / 2

	if avg > 6 {
		return true
	} else if avg < 6 {
		return false
	}

	panic("THE AVERAGE IS 6!")
}

func main() {
	fmt.Println(isStudentApproved(6, 6))
	fmt.Println("END!")
}
