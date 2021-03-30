package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorverte := trab1 || trab2

	// estudo de caso, se conseguir uma trabalha na terça chamado de trab1 e uma na quarta chamado de trab2
	// no final de semana ira comprar uma tv de 50 e tomara uma sorvete
	// se conseguir trabalho em um dos dois dias, comprar uma tv de 32 e tomar sorvete
	// se nao conseguir trabalhar em dia nenhum nao comprar nenhuma tv e nao tomar sorvete
	// retorno em forma de booleano verdadeiro ou falso

	return comprarTv50, comprarTv32, comprarSorverte
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)

}
