package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Begin Goroutines:", runtime.NumGoroutine())
	var count int

	var wg sync.WaitGroup
	wg.Add(10)

	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			val := count
			//runtime.Gosched()
			val++
			count = val
			mu.Unlock()
			wg.Done()

			fmt.Println("\n Midle Goroutines:", runtime.NumGoroutine())
			fmt.Println(count)
		}()
	}
	wg.Wait()

	fmt.Println("\nEnd Goroutines:", runtime.NumGoroutine())
	fmt.Println(count)
}