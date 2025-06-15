package main

import "fmt"

func main() {
	channel := make(chan string, 200)
	channel <- "Hello world!"
	channel <- "Coding!"
	channel <- "Coding in go!"

	message := <-channel
	message2 := <-channel

	fmt.Println(message)
	fmt.Println(message2)
}
