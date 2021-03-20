package main

import "fmt"

type person struct {
	fristName string
	lastName  string
	iceCream  []string
}

func main() {
	x := map[string]person{
		"Paula": {
			fristName: "Rodrigo",
			lastName:  "Paula",
			iceCream: []string{
				"vanilla",
				"Strawberry",
			},
		},
	}

	for i, v := range x {
		fmt.Printf("Key: %v, Value: %v", i, v)
	}

}
