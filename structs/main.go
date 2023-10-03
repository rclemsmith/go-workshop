package main

import "fmt"

type Point struct{
	x int32
	y int32
}

type Circle struct{
	radius float64
	center *Point
}

type Circle1 struct{
	radius float64
	*Point
}

func changeX(pt *Point){
	pt.x = 100
}

func main(){
	var p1 Point = Point{1,2}
	p2 := Point{-5,7}
	p3 := Point{x:2}
	p1.x = 7
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p1)
	changeX(&p1)
	fmt.Println(p1)

	c1 := Circle{4.56,&p1}

	fmt.Println(c1)
	fmt.Println(c1.center)
	fmt.Println(c1.center.x)

	c2 := Circle1{4,&p2}
	fmt.Println(c2.x)
}