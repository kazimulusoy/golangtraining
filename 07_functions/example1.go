package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	a := 41

	return a
}

func bar() (int, string) {
	var x int
	var y string

	x = 42
	y = "Hello"

	return x, y
}
