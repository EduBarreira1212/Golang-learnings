package main

import "fmt"

type user struct {
	name    string
	age     int8
	address address
}

type address struct {
	street string
	number int8
}

func main() {
	var u user
	u.name = "Eduardo"
	u.age = 21

	address1 := address{"Rua exemplo", 100}

	user2 := user{"Fulano", 50, address1}

	user3 := user{age: 30}

	fmt.Println(u)
	fmt.Println(user2)
	fmt.Println(user3)
}
