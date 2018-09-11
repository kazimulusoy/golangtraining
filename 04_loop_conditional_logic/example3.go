package main

import (
	"fmt"
	"time"
)

func main() {
	bd := 1987
	year := time.Now().Year()

	for bd <= year {
		fmt.Println(bd)
		bd++
	}
}
