package main

import (
	"12_Writing_Documentation/ex1/dog"
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
}
