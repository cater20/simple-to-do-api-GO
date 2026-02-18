package main

import (
	"fmt"
	"log"
	"net/http"

	
	"github.com/cater20/go-todo-api/routes"
	"github.com/cater20/go-todo-api/storage"
)

func main() {

	storage.InitDB() // initialize SQLite
	
	router := routes.SetupRoutes() // Get router from routes package
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
