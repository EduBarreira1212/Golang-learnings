package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	user1JSON := `{"name":"Edu","email":"edu@gmail.com","password":"12345"}`

	var user1 User

	if err := json.Unmarshal([]byte(user1JSON), &user1); err != nil {
		log.Fatal(err)
	}

	fmt.Println(user1)

	user2JSON := `{"Email":"fulano@gmail.com","Name":"Fulano","Password":"12345"}`

	user2 := make(map[string]string)

	if err := json.Unmarshal([]byte(user2JSON), &user2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(user2)
}
