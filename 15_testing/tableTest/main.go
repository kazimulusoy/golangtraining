package main

import (
	"fmt"
)

func main() {
	fmt.Println(Average(2, 4, 6, 8))
}

//Average function takes argument from slice and find average of them
func Average(num ...float64) float64 {
	fmt.Println(num, " ")
	var total float64
	var counter int

	for _, values := range num {
		total += values
		counter++
	}

	return total / float64(counter)
}
