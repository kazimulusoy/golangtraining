package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	c := make(chan int, 1)
	c <- 42
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(<-c)
}
