package main

import "fmt"

func main() {
	var b byte = 3 // atribuicao simples
	fmt.Println(b)

	i := 3 // inferencia de tipo
	i += 3 // mesma coisa de fazer i = i + 3
	i -= 3 // mesma coisa de fazer i = i - 3
	i /= 2 // mesma coisa de fazer i = i / 2
	i *= 2 // mesma coisa de fazer i = i * 2
	i %= 2 // mesma coisa de fazer i = i % 2

	fmt.Println(i)

	x, y := 1, 2 // desse modo estamos atribuindo e inizializando a variavel x recebendo 1 e variavel y recebendo 2

	fmt.Println(x, y)

	x, y = y, x // dessa forma mudamos os valores principais das variaveis as invertendo como o x e y ja existe usandos "="
	fmt.Println(x, y)

}
