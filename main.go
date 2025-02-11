package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rest-api/internal/config"
)

type Response struct {
	Message string `json:message`
}

// func handleHello(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
// 	}

// 	nome := r.URL.Query().Get("nome")

// 	if nome == "" {
// 		nome = "World"
// 	}

// 	response := Response{Message: "Hello, " + nome + "!"}

// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(response)
// }

func handleHelloGinVersion(c *gin.Context) {

	c.JSON(200, gin.H{"message": "pong"})

	// if r.Method != http.MethodGet {
	// 	http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	// }

	// nome := r.URL.Query().Get("nome")

	// if nome == "" {
	// 	nome = "World"
	// }

	// response := Response{Message: "Hello, " + nome + "!"}

	// w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder(w).Encode(response)
}

func handleTwoPointer(c *gin.Context) {

	c.JSON(200, gin.H{"message": "pong"})

}

func checkPalindromo(word string) bool {

	// arr1 := []byte(string(word))
	// runes := []rune(word)

	// for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	// 	runes[i], runes[j] = runes[j], runes[i]
	// }

	// arr2 := []byte(string(runes))

	// return reflect.DeepEqual(arr1, arr2)

	runes := []rune(word)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Connect to database
	db, err := cfg.ConnectDatabase()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Get underlying SQL DB object
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting database instance:", err)
	}

	// Test connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	log.Println("Successfully connected to database!")

	// Close connection when done
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Error closing database connection:", err)
		}
	}()
}
