package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	a := foo(xi...)
	b := bar(xi)

	fmt.Println(a)
	fmt.Println(b)
}

func foo(i ...int) int {
	var sum int
	for _, v := range i {
		sum += v
	}
	return sum
}

func bar(i []int) int {
	var sum int
	for _, v := range i {
		sum += v
	}
	return sum
}
