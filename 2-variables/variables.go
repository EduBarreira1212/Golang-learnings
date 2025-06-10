package main

import "fmt"

func main() {
	var name string = "Eduardo"
	var lastName string = "Aguiar"
	age := 21

	fmt.Println(name)
	fmt.Println(age)

	var (
		var1 int = 1
		var2 int = 2
	)

	fmt.Println(var1, var2)

	const middleName string = "Barreira"
	fmt.Println(middleName)

	name, lastName = lastName, name

	fmt.Println(name, lastName)
}
