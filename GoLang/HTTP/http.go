package main

import (
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Raiz!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Olá Mundo!</h1>"))
}

func usuários(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

	// Request - Response

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/", root)

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuários)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
