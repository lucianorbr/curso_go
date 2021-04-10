package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10000 { // não existe while em Go
		fmt.Println(i)
		time.Sleep(time.Second * 5)
		i++
	}

	for i := 0; i <= 2000000; i += 1 {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second * 5)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
			time.Sleep(time.Second * 5)
		} else {
			fmt.Print("Impar ")
			time.Sleep(time.Second * 5)
		}
	}

	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 5)
	}

}
