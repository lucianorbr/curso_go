package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas forma posfix
	x++ // pode se usar x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // pode se usar y -= 1 ou y = y -1
	fmt.Println(y)

}
