package main

import "fmt"

type veículo struct {
	name  string
	cor   string
	porta int
}

type caminhonete struct {
	veículo         veículo
	traçãoNasQuatro bool
}

type sedan struct {
	veículo    veículo
	modeloLuxo bool
}

func main() {
	duster := caminhonete{
		veículo: veículo{
			name:  "duster",
			cor:   "blak",
			porta: 4,
		},
		traçãoNasQuatro: false,
	}

	fiesta := sedan{
		veículo: veículo{
			name:  "fiesta",
			cor:   "blue",
			porta: 4,
		},
		modeloLuxo: false,
	}

	fmt.Printf("%v\n", duster)

	fmt.Printf("%v", fiesta)
}
