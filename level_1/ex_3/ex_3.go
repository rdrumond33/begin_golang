package main

import "fmt"

var (
	x = 42
	y = "James Bond"
	z = true
)

func main() {
	s := fmt.Sprintf("Values: x(%v: %T), y(%v: %T) , z(%v: %T)", x, x, y, y, z, z)
	fmt.Print(s)
}
