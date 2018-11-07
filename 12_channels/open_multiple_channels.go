package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	fmt.Println("ROUTINES", runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}

	fmt.Println("ROUTINES", runtime.NumGoroutine())
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("About the exit...")
}
