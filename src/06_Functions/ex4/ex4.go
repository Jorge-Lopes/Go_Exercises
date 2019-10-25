package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last)
}
func main() {
	p := person{
		first: "Jorge",
		last:  "Lopes",
		age:   24,
	}
	p.speak()
}
