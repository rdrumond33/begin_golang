package main

import "fmt"

type pessoa struct {
	name  string
	idade int
}

func main() {
	x := pessoa{
		name:  "Rodrigo",
		idade: 50,
	}

	fmt.Println("nao mudou", x)
	mudar(&x)
	fmt.Println("Mudou", x)
}

func mudar(p *pessoa) {
	(*p).name = "Mudar"
	p.idade = 5
}
