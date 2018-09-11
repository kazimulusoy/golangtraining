package main

import (
	"fmt"
)

func main() {
	mp := make(map[string][]string)

	mp[`james_bond`] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	mp[`moneypenny_miss`] = []string{`James Bond`, `Literature`, `Computer Science`}
	mp[`no_dr`] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for i, v := range mp {
		fmt.Printf("%v\n", i)
		for _, v2 := range v {
			fmt.Printf("\t%v ", v2)
		}
		fmt.Println()
	}
}
