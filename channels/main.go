package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int,2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch <-chan int , wg *sync.WaitGroup){
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh,wg)

	go func(ch chan<- int , wg *sync.WaitGroup){
		// close(myCh)
		myCh <- 5
		// myCh <- 6
		wg.Done()
	}(myCh,wg)

	// myCh <- 5 
	// fmt.Println(<-myCh)

	wg.Wait()
}
