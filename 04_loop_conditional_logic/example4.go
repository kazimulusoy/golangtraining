package main

import (
	"fmt"
	"time"
)

func main() {
	bd := 1987
	year := time.Now().Year()

	for {
		if bd > year {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
