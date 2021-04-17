package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade string
}

func (p *pessoa) falar() {
	fmt.Println("Ola !!!")
	fmt.Println("Nome:", (*p).nome)
	fmt.Println("Idade:", (*p).idade)
}

type humano interface {
	falar()
}

func main() {
	p := pessoa{
		nome:  "Rodrigo",
		idade: "50",
	}

	dizerAlgumaCoisa(&p)
}

func dizerAlgumaCoisa(h humano) {
	h.falar()
}
