package main

import (
	"fmt"
)

func main() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	sl1 := append(sl[:5])
	sl2 := append(sl[5:])
	sl3 := append(sl[2:7])
	sl4 := append(sl[1:6])

	fmt.Println(sl1)
	fmt.Println(sl2)
	fmt.Println(sl3)
	fmt.Println(sl4)
}
