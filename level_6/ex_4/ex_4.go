package main

import "fmt"

type pessoa struct {
	name      string
	sobrenome string
	idade     int
}

func (p pessoa) imprimir() string {
	return fmt.Sprintf("Nome: %v, Sobrenome: %v, Idade: %d", p.name, p.sobrenome, p.idade)
}
func main() {
	p := pessoa{
		name:      "Rodrigo",
		sobrenome: "Paula",
		idade:     50,
	}

	fmt.Println(p.imprimir())
}
