package main

import "fmt"

type valuex int

var x valuex

func main() {

	fmt.Printf("x(%v: %T)", x, x)
	x = 42
	fmt.Printf("\nx(%v: %T)", x, x)

}
