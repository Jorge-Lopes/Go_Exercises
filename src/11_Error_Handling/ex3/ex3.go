package main

import "fmt"

type customErr struct {
	info string
}

type error interface {
	Error() string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Error found at: %v", ce.info)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	c := customErr{
		info: "Custom Error",
	}
	foo(c)

}
