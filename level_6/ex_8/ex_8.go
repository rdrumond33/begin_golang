package main

import "fmt"

func main() {
	x := soma(1, 2)
	fmt.Println(x())
}

func soma(x, y int) func() string {
	return func() string {
		return fmt.Sprintf("A soma e: %v", (x + y))
	}
}
