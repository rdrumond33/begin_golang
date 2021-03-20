package main

import "fmt"

func main() {
	x := struct {
		name        string
		cor         []string
		equipamento map[string][]string
	}{
		name: "Rodrigo",
		cor: []string{
			"blue",
			"blak",
			"orange",
		},
		equipamento: map[string][]string{
			"notebbok": {
				"acer",
				"positivo",
			},
			"televisao": {
				"acer",
				"sansumg",
			},
		},
	}

	fmt.Println(x)
}
