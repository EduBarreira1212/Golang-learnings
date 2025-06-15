package main

import (
	"fmt"
	"time"
)

func main() {
	canal := write("Hello world!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
