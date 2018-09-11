package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sl)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", sl)
}
