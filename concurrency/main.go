package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"Test"}

var wg sync.WaitGroup // Pointer

var mut sync.Mutex 

func main() {
	// go greeter("Hello")
	// go greeter("World")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
	}

	for _,web := range websitelist{
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string){
	 
	res,err := http.Get(endpoint)
	
	if err != nil {
		fmt.Println("OOPS in Endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d Status Code for website for %s\n", res.StatusCode,endpoint)
	}

	wg.Done()
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}
