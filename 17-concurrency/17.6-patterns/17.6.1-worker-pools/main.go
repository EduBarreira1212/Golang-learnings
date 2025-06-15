package main

import "fmt"

func main() {
	works := make(chan int, 45)
	results := make(chan int, 45)

	go worker(works, results)
	go worker(works, results)
	go worker(works, results)
	go worker(works, results)

	for i := 0; i < 45; i++ {
		works <- i
	}
	close(works)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(works <-chan int, results chan<- int) {
	for number := range works {
		results <- fibonacci(number)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
