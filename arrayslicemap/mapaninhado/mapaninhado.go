package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 1144.56,
		},
		"L": {
			"Luciano": 252454.56,
			"Leliana": 24554.55,
		},
		"J": {
			"João": 55455.56,
		},
	}

	//delete(funcsPorLetra, "J")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

}
