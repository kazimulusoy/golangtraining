package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p := person{
		name: "James Bond",
	}

	fmt.Println(p.name)
	changeMe(&p)
	fmt.Println(p.name)
}

func changeMe(person *person) {
	//person.name = "Lara Croft"
	(*person).name = "Lara Croft"
}
