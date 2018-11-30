package main

import (
	"fmt"

	"github.com/kazimulusoy/golangtraining/15_testing/BET_example2/quote"
	"github.com/kazimulusoy/golangtraining/15_testing/BET_example2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
