package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Hello from anonymous func.")
	}()

}

// func hfaf() {
// 	fmt.Println("Hello from anonymous func.")
// }
