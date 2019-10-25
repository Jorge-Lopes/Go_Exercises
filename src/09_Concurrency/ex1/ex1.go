package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("\nBegin Goroutine", runtime.NumGoroutine())
	fmt.Println("Begin CPU", runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("\nGoroutine 2")
		wg.Done()

	}()

	fmt.Println("\nMidle Goroutine", runtime.NumGoroutine())
	fmt.Println("Midle CPU", runtime.NumCPU())

	wg.Wait()
	fmt.Println("Goroutine main")
	fmt.Println("\nEnd Goroutine", runtime.NumGoroutine())
	fmt.Println("End CPU", runtime.NumCPU())
}
