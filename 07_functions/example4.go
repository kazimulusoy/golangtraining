package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello from " + p.first + " " + p.last)
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   31,
	}

	p.speak()
}
