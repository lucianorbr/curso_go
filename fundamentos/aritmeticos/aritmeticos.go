package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// operacoes bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// outras operacoes usando math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b))) // compara e retorno quem é o maior valor entra as variaveis a e b, foi preciso converver pois a funcao apenas aceita float
	fmt.Println("Menor =>", math.Min(c, d))                   // compra e retorna quem é o menor valor das variaveis c e d, não precisa converter pois já sao float
	fmt.Println("Exponenciação =>", math.Pow(c, d))           // funcao volta um valor elevado ao outro,

}
