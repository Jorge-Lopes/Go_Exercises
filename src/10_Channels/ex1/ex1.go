package main

import "fmt"

func main() {
	// func literal
	c := make(chan int)
	go func() {
		c <- 42
	}()

	//buffered channel
	//c := make(chan int, 1)
	//c <- 42

	fmt.Println(<-c)
}
