package main

import "fmt"

func main() {
	xi := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for i, v := range xi {
		fmt.Println(i, v)
	}
	fmt.Println(xi[2:7])
	fmt.Println(xi[7:])
	fmt.Println(xi[4:9])
	fmt.Println(xi[:6])
}