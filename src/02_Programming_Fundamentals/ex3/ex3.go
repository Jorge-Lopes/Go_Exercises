package main

import "fmt"

const (
	x     = 1
	y int = 2
)

func main() {
	fmt.Printf("%v of type %T\n", x, x)
	fmt.Printf("%v of type %T\n", y, y)
}
