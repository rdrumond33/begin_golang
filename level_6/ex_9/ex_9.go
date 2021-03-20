package main

import "fmt"

func main() {
	x := par(soma, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Soma Par", x)
	y := par(mult, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("subtração Par", y)

}

func soma(x ...int) int {
	res := 0
	for _, v := range x {
		res = res + v
	}
	return res
}

func mult(x ...int) int {
	res := 1
	for _, v := range x {
		res = res * v
	}
	return res
}
func par(calback func(x ...int) int, x ...int) int {
	var pares []int
	for _, v := range x {
		if v%2 == 0 {
			pares = append(pares, v)
		}
	}
	return calback(pares...)
}
