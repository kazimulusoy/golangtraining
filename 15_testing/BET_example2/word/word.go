package word

import (
	"fmt"
	"reflect"
	"strings"
)

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//this function is which for return for string count
func Count(s string) int {
	l := strings.Fields(s)
	fmt.Println(reflect.TypeOf(l))
	return len(l)
}
