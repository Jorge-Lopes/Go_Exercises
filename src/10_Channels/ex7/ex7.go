package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		c := make(chan int)
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			close(c)
		}()

		for v := range c {
			fmt.Println(v)

		}
		fmt.Println("-----")
	}

	fmt.Println("End")

}
