package main

import (
	"fmt"
)

type person struct {
	name    string
	surname string
	age     int
}

type human interface {
	personPrinter() string
}

func (p *person) personPrinter() string {
	info := fmt.Sprint(p.name, p.surname, p.age)

	return info
}

func saySomething(h human) {
	fmt.Println("Human is: " + h.personPrinter())
}

func main() {
	p := person{
		name:    "James",
		surname: "Bond",
		age:     32,
	}

	//saySomething(p)
	saySomething(&p)
	fmt.Println(p.personPrinter())
}
