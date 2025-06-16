package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(write("Hello!"), write("Golang"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(channel1, channel2 <-chan string) <-chan string {
	exitChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-channel1:
				exitChannel <- message
			case message := <-channel2:
				exitChannel <- message
			}
		}
	}()

	return exitChannel

}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
