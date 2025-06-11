package main

import "fmt"

func main() {
	person := map[string]string{
		"name":     "Eduardo",
		"lastName": "Aguiar",
	}

	fmt.Println(person)
	fmt.Println(person["name"])

	person2 := map[string]map[string]string{
		"name": {
			"first": "Fulano",
			"last":  "Da Silva",
		},
		"course": {
			"name":   "Medicine",
			"campus": "Campus 1",
		},
	}

	fmt.Println(person2)
	delete(person2, "name")
	fmt.Println(person2)

	person2["zodiacSign"] = map[string]string{
		"name": "sagittarius",
	}

	fmt.Println(person2)
}
