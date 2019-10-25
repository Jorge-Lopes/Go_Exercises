package main

import "fmt"

func main() {
	s := foo()
	fmt.Println(s())
}

func foo() func() string {
	return func() string {
		return fmt.Sprint("Return function")
	}
}
