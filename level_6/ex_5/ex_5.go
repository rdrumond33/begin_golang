package main

import (
	"fmt"
	"math"
)

type figura interface {
	area() float64
}

type quadrado struct {
	lado float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

func info(x figura) string {
	return fmt.Sprintf("%f", x.area())
}

func main() {
	quadrado := quadrado{
		lado: 10,
	}

	circulo := circulo{
		raio: 15.5,
	}

	fmt.Println("Vaolora da area do quadrado", info(quadrado))

	fmt.Println("Vaolora da area do quadrado", info(circulo))
}
