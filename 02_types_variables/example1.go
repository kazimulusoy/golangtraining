package main

import (
	"fmt"
)

var number int

func main() {
	number = 47

	fmt.Printf("%d, %b, %#x", number, number, number)
}
