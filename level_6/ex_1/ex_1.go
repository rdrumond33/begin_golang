package main

import "fmt"

func main() {
	fmt.Println(soma(1, 2))
	i, s := ola()
	fmt.Println(i, s)
}

func soma(x, y int) int {
	return x + y
}

func ola() (int, string) {
	return 1, "ola"
}
