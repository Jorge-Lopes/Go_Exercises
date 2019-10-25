package main

import "fmt"

func main() {
	xi := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for i, v := range xi {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", xi)
}
