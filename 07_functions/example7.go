package main

import (
	"fmt"
)

var index int
var f2 func()

func main() {
	fourtyTwo := func() {
		for index = 0; index < 3; index++ {
			fmt.Println(index)
		}
	}

	fourtyTwo()
	fmt.Printf("%T\n", fourtyTwo)

	f2 = fourtyTwo
	f2()
	fmt.Printf("%T\n", f2)
}
