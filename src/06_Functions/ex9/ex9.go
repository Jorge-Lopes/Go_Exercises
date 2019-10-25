package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b := func(xi []int) int {
		var i int
		for _, v := range xi {
			i = i + v
		}
		return i
	}
	fmt.Println(b(xi))
	x := foo(b, xi)
	fmt.Println(x)
}

func foo(f func(xi []int) int, ii []int) int {
	x := f(ii)
	x++
	return x
}
