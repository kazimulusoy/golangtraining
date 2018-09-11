package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"hello":       []string{"hello", "world"},
		"programming": []string{"programming", "example"},
		"country":     []string{"Germany", "France", "Italy"},
	}

	fmt.Println(m)

	delete(m, "country")

	fmt.Println(m)
}
