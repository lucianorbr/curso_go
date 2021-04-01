package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s)) // len pega o tamanho do item desejado no nosso caso s, cap pega a capacidade do item desejado no nosso caso s
	// append acresta itens no nosso caso dentro do s que é nosso slice

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) // qdo atinge a capacidade maximo em um slice que no nosso caso é 20 ele automaticamente aumenta seu proprio tamanho
}
