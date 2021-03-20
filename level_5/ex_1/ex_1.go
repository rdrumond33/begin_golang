package main

import "fmt"

type person struct {
	fristName string
	lastName  string
	iceCream  []string
}

func main() {
	x := person{
		fristName: "Rodrigo",
		lastName:  "Paula",
		iceCream: []string{
			"vanilla",
			"Strawberry",
		},
	}
	fmt.Println(x)

}
