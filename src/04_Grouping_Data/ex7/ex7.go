package main

import "fmt"

func main() {
	x := [][]string{
		{"james", "bond", "shaken"},
		{"miss", "money", "hello"},
	}
	for i, v := range x {
		fmt.Println(i, v)
		for j, p := range v {
			fmt.Println(j, p)
		}
	}

}
