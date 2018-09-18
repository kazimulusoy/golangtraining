package main

import (
	"fmt"
)

type person struct {
	firstName      string
	lastName       string
	iceCreamFlavor []string
}

func main() {
	p1 := person{
		firstName:      "James",
		lastName:       "Bond",
		iceCreamFlavor: []string{"vanilla", "banana", "lemon"},
	}

	p2 := person{
		firstName:      "Jack",
		lastName:       "London",
		iceCreamFlavor: []string{"strawberry", "coconut"},
	}

	fmt.Println(p1, p2)

	fmt.Println("p1 favorite ice creams:")
	for i, v := range p1.iceCreamFlavor {
		fmt.Println(i, v)
	}

	fmt.Println("p2 favorite ice creams:")
	for i, v := range p2.iceCreamFlavor {
		fmt.Println(i, v)
	}
}
