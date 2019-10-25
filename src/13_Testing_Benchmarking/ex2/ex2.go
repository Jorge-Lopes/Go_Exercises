package main

import (
	"13_Testing_Benchmarking/ex2/quote"
	"13_Testing_Benchmarking/ex2/word"
	"fmt"
)

func main() {

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

	fmt.Println(word.Count(quote.SunAlso))
}
