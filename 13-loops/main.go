package main

import (
	"fmt"
)

func main() {
	i := 0

	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("------------------")

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	fmt.Println("------------------")

	names := []string{"Fulano", "Beltrano", "Ciclano"}

	for i, name := range names {
		fmt.Println(i, name)
	}

	fmt.Println("------------------")

	for i, letter := range "WORD" {
		fmt.Println(i, string(letter))
	}

	fmt.Println("------------------")

	user := map[string]string{
		"name":     "Fulano",
		"lastName": "Silva",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}
}
