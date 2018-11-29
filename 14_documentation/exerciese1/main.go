package main

import (
	"fmt"

	"github.com/kazimulusoy/golangtraining/14_documentation/exerciese1/dog"
)

type canine struct {
	name string
	age  int
}

//godoc --http=:8080 under package documantation there will be dog package documentation
func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(3),
	}

	fmt.Println(fido)
}
