package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "coffee time",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
