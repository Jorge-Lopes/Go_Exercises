package main

import "fmt"

func main() {
	x := 1
	y := 2

	g := x == y
	h := x <= y
	i := x >= y
	j := x != y
	k := x < y
	l := x > y

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}
