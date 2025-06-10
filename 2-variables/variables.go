package main

import "fmt"

func main() {
	var name string = "Eduardo"
	age := 21

	fmt.Println(name)
	fmt.Println(age)

	var (
		var1 int = 1
		var2 int = 2
	)

	fmt.Println(var1, var2)

	const lastName string = "Barreira"
	fmt.Println(lastName)
}
