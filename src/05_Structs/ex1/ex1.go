package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	p1 := person{
		firstName: "Jorge",
		lastName:  "Lopes",
		iceCream:  []string{"Chocolate"},
	}

	p2 := person{
		firstName: "Jianyu",
		lastName:  "Li",
		iceCream:  []string{"Vanilla"},
	}

	fmt.Println(p1.firstName, p1.lastName, p1.iceCream)
	fmt.Println(p2.firstName, p2.lastName, p2.iceCream)
}
