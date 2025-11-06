package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP é um protocolo de comunicação - Base da comunicalçao WEB

	// Cliente (faz requisisçao) - Servidor (processa requiseção e envia resposta)

	// Request - response

	// Rotas
	// URI - Identificador do recurso
	// Metodo - GET, POST, PUT, DELETE

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pagina Raiz"))
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ola Mundo!"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregador de pagina de usuarios"))
	})

	log.Fatal(http.ListenAndServe(":500", nil))
}
