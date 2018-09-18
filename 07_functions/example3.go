package main

import (
	"fmt"
)

func main() {

	defer h()
	w()
	h()
	w()
}

func h() {
	fmt.Println("Hello")
}

func w() {
	fmt.Println("World")
}
