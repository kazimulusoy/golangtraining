package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {
	c := make(chan int)

	//send
	//go foo(c)
	go func() {
		mutex.Lock()
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		mutex.Unlock()
	}()

	//receive
	//bar(c)
	func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	fmt.Println("Program exit...")
}

// send channel
// func foo(c chan<- int) {
// 	for i := 0; i < 100; i++ {
// 		c <- i
// 	}
// 	close(c)
// }

//receive channel
// func bar(c <-chan int) {
// 	fmt.Println(<-c)
// }
