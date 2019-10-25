package main

import "fmt"

func main() {
	t := struct {
		doors     int
		color     string
		fourWheel bool
	}{

		doors:     4,
		color:     "blue",
		fourWheel: true,
	}

	fmt.Println(t)
}
