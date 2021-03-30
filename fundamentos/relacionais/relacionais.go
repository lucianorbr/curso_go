package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana") // operador de igualdade
	fmt.Println("!=", 3 != 2)                     // operador diferente de ....
	fmt.Println("<", 3 < 2)                       // operador menor que .....
	fmt.Println(">", 3 > 2)                       // operador maior que .....
	fmt.Println("<=", 3 <= 2)                     // operador menor ou igual a ....
	fmt.Println(">=", 3 >= 2)                     // operador maior ou igual a ....

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d2.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{Nome: "Joã"}
	fmt.Println("Pessoas:", p1 == p2)

}
