package hello

import (
	"fmt"
	"testing"
)

//test func
func TestHello(t *testing.T) {
	s := Hello("Welt!")

	if s != "Hello Welt!" {
		t.Error("got " + s + " expected " + "Hello Welt!")
	}
}

func ExampleHello() {
	fmt.Println(Hello("World"))
	//Output:
	//Hello World
}

//benchmark func
func BenchmarkHello(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Hello("World!")
	}
}

/*
When you try this command "go test -bench ." expected output will be:
BenchmarkHello-8        100000000               21.0 ns/op
PASS
That means from 8 cores, 100.000.000 times called the function and each operation
has 19.5ns approx operation time.
*/
/*
Here is some test commands for golang code:
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out

go tool
go tool cover -h
*/
