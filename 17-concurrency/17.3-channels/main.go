package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Hello world", channel)

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("END!")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
