package main

import "fmt"

// - Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
// - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
// - Do quinto ao último item do slice (incluindo o último item!)
// - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
// - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
// - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

func main() {
	x := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("1- %v\n", x[:3])
	fmt.Printf("2- %v\n", x[5:])
	fmt.Printf("3- %v\n", x[1:8])
	fmt.Printf("4- %v\n", x[2:9])
	fmt.Printf("5- %v\n", x[2:len(x)-1])

}
