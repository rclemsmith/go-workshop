package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32{
	sum := 0 
	for _, v := range s.grades {
		sum += v
	}

	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int{
	curMax :=0 
	for _,v := range s.grades{
		if curMax < v {
			curMax = v
		}
	}

	return curMax
}

func main() {
	s1 := Student{"Tim", []int{70, 80, 90, 85}, 19}
	fmt.Println(s1.getAge())
	s1.setAge(7)
	fmt.Println(s1.getAge())

	fmt.Println(s1.getAverageGrade())

	fmt.Println(s1.getMaxGrade())
}
