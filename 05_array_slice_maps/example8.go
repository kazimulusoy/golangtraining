package main

import (
	"fmt"
)

func main() {
	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(s1)
	fmt.Println(s2)

	ss := [][]string{s1, s2}
	fmt.Println(ss)

	for i := 0; i < len(ss); i++ {
		for j := 0; j < len(ss[0]); j++ {
			fmt.Println(ss[i][j])
		}
	}

}
