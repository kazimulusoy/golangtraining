package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	tr := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	sd := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "grey",
		},
		luxury: true,
	}

	fmt.Println(tr, sd)
	fmt.Println(tr.doors)
	fmt.Println(sd.doors)

}
