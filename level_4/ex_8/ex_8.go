package main

import "fmt"

func main() {
	x := make(map[string][]string)

	x["Paula"] = []string{
		"comer",
		"correr",
		"Programar",
	}

	x["Drumond"] = []string{
		"Programar",
	}

	for i, v := range x {
		fmt.Printf("Sobrenome: %v, Hob: %v\n", i, v)
	}
}
