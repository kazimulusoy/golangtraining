package main

import (
	"fmt"
)

func main() {

	f := foo()
	fmt.Println(f())
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", f())

}

func foo() func() int {
	return func() int {
		return 42
	}
}
