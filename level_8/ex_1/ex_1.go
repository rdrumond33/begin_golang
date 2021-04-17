package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	First string `json:"frist"`
	Age   int    `json:"age"`
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	userJSON, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON: %s", userJSON)
}
