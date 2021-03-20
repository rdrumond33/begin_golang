package main

import "fmt"

type valuex int

type valuey valuex

var (
	x valuex
	y int
)

func main() {
	fmt.Printf("x(%v:%T),y(%v:%T)", x, x, y, y)
	x = 42
	y = int(x)
	fmt.Printf("\nx(%v:%T),y(%v:%T)", x, x, y, y)

}
