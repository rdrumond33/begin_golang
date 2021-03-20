package main

import "fmt"

func main() {
	x := func(x, y int) int {
		return x + y
	}

	fmt.Println(x(1, 2))
}
