package main

import "fmt"

func main() {
	i := 1
	if i == 10 {
		fmt.Println("Correct")
	} else if i > 10 {
		fmt.Println("Too high")
	} else {
		fmt.Println("Too low")
	}
}
