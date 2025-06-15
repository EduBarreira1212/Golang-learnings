package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello world!") // goroutine
	write("Coding in Go!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
