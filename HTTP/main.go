package main // Alterado de http para main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola mundo"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar pagina de usuarios"))
}

func main() {
	// Registrando as rotas corretamente
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	// Iniciando o servidor na porta 5000
	log.Println("Servidor rodando na porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
