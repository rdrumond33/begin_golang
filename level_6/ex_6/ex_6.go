package main

import "fmt"

func main() {
	func(x, y int) {
		fmt.Println("soma", x+y)
	}(13, 22)
}
