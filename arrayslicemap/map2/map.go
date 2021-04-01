package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José":  1144.55,
		"Maria": 144455.22,
		"Pedro": 14455.14,
	}

	funcsESalarios["Luciano"] = 5546.55
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
