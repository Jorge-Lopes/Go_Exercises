package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() string {
	s := fmt.Sprint("I am ", p.name, " and i have ", p.age)
	return s
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	p := person{
		name: "Jorge",
		age:  24,
	}

	saySomething(&p)
	p.speak()
	//saySomething(p)
}
