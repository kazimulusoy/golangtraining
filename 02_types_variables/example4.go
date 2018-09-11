package main

import (
	"fmt"
)

var num int

func main() {
	num = 8
	fmt.Printf("decimal:%d, binary:%b, hex:%#x\n", num, num, num)

	num = num << 1
	fmt.Printf("decimal:%d, binary:%b, hex:%#x", num, num, num)
}
