package main

import "fmt"

func main() {
	i := 1

	// Go não te aritmética de ponteiros
	var p *int = nil

	p = &i // pegando o endereço da variavel i e atribuindo ao um ponteiro
	*p++
	i++

	fmt.Println(*p, p, i)
}
