package main

import (
	"fmt"
)

func main() {

	fmt.Println("Factorial example with recursion.")
	n := 10

	fmt.Println(factorial(n))
}

func factorial(n int) int {
	if n < 2 {
		return 1
	}

	return n * factorial(n-1)
}

// Compexity is O(n) because recursive function linearly passing through the numbers
// 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
