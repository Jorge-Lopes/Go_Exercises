package main

import "fmt"

var r [5]int

func main() {
	for i := 0; i < len(r); i++ {
		r[i] = i + 10
	}
	for i, v := range r {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", r)
}
