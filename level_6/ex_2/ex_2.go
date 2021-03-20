package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	fmt.Println("Usando sem slice", soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13))
	fmt.Println("Usando slice", soma(x...))

	fmt.Println("Usando soma2", soma2(x))

}

func soma(x ...int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func soma2(x []int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
