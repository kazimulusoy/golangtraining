package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"first" `
	Age   int    `json:"age"`
	Type  string `json:"type"`
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
		Type:  "M",
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
		Type:  "M",
	}

	u3 := user{
		First: "M",
		Age:   54,
		Type:  "Not defined",
	}

	users := []user{u1, u2, u3}

	json, err := json.Marshal(users)

	if err != nil {
		fmt.Println("json marshal error")
	}

	fmt.Println(json)
	fmt.Println(string(json))
}
