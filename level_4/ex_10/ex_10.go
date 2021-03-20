package main

import "fmt"

func main() {
	x := map[string][]string{
		"Paula": {
			"Correr",
			"Limpar",
		},
		"Drumond": {
			"Nadar",
			"Programar",
		},
	}

	delete(x, "Paula")

	for i, v := range x {
		fmt.Printf("Index: %v, Hobb: %v", i, v)
	}
}
