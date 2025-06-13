package main

import "fmt"

func main() {
	result := func(text string) string {
		return fmt.Sprintf("Text => %s", text)
	}("Test")

	fmt.Println(result)
}
