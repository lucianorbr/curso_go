package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public")) // criar uma variavel que armazendo a criacao do servidor, apontado para a pasta public
	http.Handle("/", fs)                      // cria um handle apontando o diretorio "/" para a index do servisor

	log.Println("Executando...") // criar um log que mostra se o servidor esta ativo
	log.Fatal(http.ListenAndServe(":3000", nil))

}
