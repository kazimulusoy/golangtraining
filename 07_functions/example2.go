package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(bar([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func foo(variadicParam ...int) int {
	total := 0

	for _, v := range variadicParam {
		total += v
	}

	return total
}

func bar(variadicParam []int) int {
	total := 0

	for _, v := range variadicParam {
		total += v
	}

	return total
}
