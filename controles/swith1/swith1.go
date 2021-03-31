package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough // funcao executa o primeiro case e sai executando os proximos para atender as condiccoes
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println("A nota é:", notaParaConceito(9.5))
	fmt.Println("A nota é:", notaParaConceito(8))
	fmt.Println("A nota é:", notaParaConceito(5))
	fmt.Println("A nota é:", notaParaConceito(3))
	fmt.Println("A nota é:", notaParaConceito(0))
}
