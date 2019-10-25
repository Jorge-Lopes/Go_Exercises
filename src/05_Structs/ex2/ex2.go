package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	var m = map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for i, v := range m {
		fmt.Println(i, v.first)
		for j, k := range v.favFlavors {
			fmt.Println(j, k)
		}
	}
}
