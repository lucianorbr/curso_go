package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	for {
		switch {
		case t.Hour() < 12:
			fmt.Println("Bom dia")
			time.Sleep(time.Second * 5)
		case t.Hour() < 19:
			fmt.Println("Boa tarde")
			time.Sleep(time.Second * 5)
		default:
			fmt.Println("Boa noite")
			time.Sleep(time.Second * 5)

		}
	}

}
