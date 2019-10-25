package main

import "fmt"

func main() {
	defer bar()
	foo()
}

func foo() {
	fmt.Println("I will start.")
}

func bar() {
	fmt.Println("I will wait and finish.")
}
