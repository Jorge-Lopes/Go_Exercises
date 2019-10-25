package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("Begin Goroutines:", runtime.NumGoroutine())
	var count int64

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println(atomic.LoadInt64(&count))
			wg.Done()

			fmt.Println("\n Midle Goroutines:", runtime.NumGoroutine())
		}()
	}
	wg.Wait()

	fmt.Println("\nEnd Goroutines:", runtime.NumGoroutine())
	fmt.Println(count)
}
