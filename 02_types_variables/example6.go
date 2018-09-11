package main

import (
	"fmt"
)

const (
	year int = 2014
)

const (
	a = year + iota
	b = year + iota
	c = year + iota
	d = year + iota
)

// this is iota example
func main() {
	fmt.Printf("a:%d, b:%d, c:%d, d:%d", a, b, c, d)
}
