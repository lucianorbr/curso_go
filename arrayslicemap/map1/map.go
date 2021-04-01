package main

import "fmt"

func main() {
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[2232332332] = "Maria"
	aprovados[2332332332] = "Maria A "
	aprovados[2134332332] = "Maria B "
	aprovados[2232332332] = "Maria D"
	aprovados[22323323325] = "Maria W"
	aprovados[2232372332] = "Maria E"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[2332332332])
	delete(aprovados, 2332332332)
	fmt.Println(aprovados[2332332332])

}
