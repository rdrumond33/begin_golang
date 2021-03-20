package main

import "fmt"

const (
	a         = 10
	b         = 20.0
	c int     = 50
	d float64 = 60
)

func main() {
	fmt.Printf("a(%v-%T),b(%v-%T),c(%v-%T),d(%v-%T)", a, a, b, b, c, c, d, d)
}
