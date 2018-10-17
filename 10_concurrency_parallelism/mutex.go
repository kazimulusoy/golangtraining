package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	counter := 0
	worker := 100

	wg.Add(worker)

	for i := 0; i < worker; i++ {
		go func() {
			mutex.Lock()
			count := counter
			count++
			counter = count
			fmt.Println(counter)
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
