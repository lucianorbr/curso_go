package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// formatacao de data em Go
// 01 = mes
// 02 = dia
// 03 = hora
// 04 = minutos
// 05 = segundos
// 06 = ano

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1> Hora certa: %s</h1>", s)

}

func main() {
	http.HandleFunc("/", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3030", nil))
}
