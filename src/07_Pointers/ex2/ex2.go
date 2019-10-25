package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	p.name = "Daniel"
	p.age = 25
}

func main() {
	p1 := person{
		name: "Jorge",
		age:  24,
	}

	fmt.Println(p1.name, p1.age)
	changeMe(&p1)
	fmt.Println(p1.name, p1.age)
}
