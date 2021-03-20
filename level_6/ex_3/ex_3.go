package main

import "fmt"

func main() {
	fmt.Println("Usando sem slice", soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13))

}

func soma(x ...int) int {
	defer fmt.Println("terminou o range")
	sum := 0
	fmt.Println("Comecou o range")
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
