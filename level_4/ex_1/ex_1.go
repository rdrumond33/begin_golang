package main

import "fmt"

func main() {
	x := [5]int{}
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[4] = 4

	for _, v := range x {
		fmt.Printf("Valor: %v, Tipo: %T\n", v, v)
	}

	fmt.Printf("Tipo do Array: %T", x)
}
