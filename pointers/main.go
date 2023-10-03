package main

import "fmt"

func changeValue(str *string){
	*str  = "changed!"
}

func changeValueTo (str string){
	str = "Changed23!"
}

func main() {
	x := 7
	y := &x
	fmt.Println(&x)
	fmt.Println(x)
	*y = 8
	fmt.Println(x)

	toChange := "Hello"
	changeValue(&toChange)
	fmt.Println(toChange)
	changeValueTo(toChange)
	fmt.Println(toChange)


	var pointer *string = &toChange
	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Println(&pointer)
}
