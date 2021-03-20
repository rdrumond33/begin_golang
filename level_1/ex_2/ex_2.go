package main

import "fmt"

var (
	x int
	y string
	z bool
)

func main() {
	fmt.Printf("Default value: %v :%T, %v :%T, %v :%T", x, x, y, y, z, z)
}
