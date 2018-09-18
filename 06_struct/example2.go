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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.iceCreamFlavor {
			fmt.Println(i, val)
		}
	}
}
