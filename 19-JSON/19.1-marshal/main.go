package main

import (
	"bytes"
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
	user1 := User{"Edu", "edu@gmail.com", "12345"}

	user1JSON, err := json.Marshal(user1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(user1JSON))

	user2 := map[string]string{
		"Name":     "Fulano",
		"Email":    "fulano@gmail.com",
		"Password": "12345",
	}

	user2JSON, err := json.Marshal(user2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(user2JSON))
}
