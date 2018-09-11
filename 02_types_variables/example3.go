package main

import (
	"fmt"
)

const (
	a     = 42
	b int = 43
)

func main() {
	fmt.Println(a, b)
	fmt.Printf("a:%T, b:%T", a, b)
}
