package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:message`
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	nome := r.URL.Query().Get("nome")

	if nome == "" {
		nome = "World"
	}

	response := Response{Message: "Hello, " + nome + "!"}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/hello", handleHello)
	log.Print("Servidor iniciado na porta 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal((err))
	}
}
