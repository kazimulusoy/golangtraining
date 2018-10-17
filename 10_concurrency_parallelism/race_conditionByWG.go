package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	counter := 0
	worker := 100

	wg.Add(worker)

	for i := 0; i < worker; i++ {
		go func() {
			runtime.Gosched()
			count := counter
			count++
			counter = count
			fmt.Println(counter)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
