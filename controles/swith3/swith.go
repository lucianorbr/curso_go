package main

import "fmt"

func tipo(i interface{}) string { // padrao generico para inferior o tipo passado
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "Não sei"

	}
}

func main() {
	fmt.Println("O tipo é:", tipo("sou eu"))
	fmt.Println("O tipo é:", tipo(1))
	fmt.Println("O tipo é:", tipo(9.5))
	fmt.Println("O tipo é:", tipo(func() {}))

}
