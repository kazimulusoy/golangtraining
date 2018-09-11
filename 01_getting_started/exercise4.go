package main

import (
	"fmt"
)

type my_type int

var x my_type

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
