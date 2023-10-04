package main

import (
	"fmt"
	"sync"
)


func main() {

	var wg sync.WaitGroup
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	wg.Add(2)
	go func(jobs chan int, results chan int) {
		for i := 0; i < 100; i++ {
			jobs <- i
		}
		wg.Done()
	}(jobs, results)

	go func(jobs chan int, results chan int) {
		for i := 0; i < 100; i++ {
			fmt.Println(<-results)
		}
		wg.Done()
	}(jobs, results)

	wg.Wait()
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 0 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
