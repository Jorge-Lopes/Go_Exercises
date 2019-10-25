package main

import (
	"13_Testing_Benchmarking/ex1/dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "fido",
		age:  dog.Years(3),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(3))
}
