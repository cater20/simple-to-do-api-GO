package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yourname/go-todo-api/routes"
)

func main() {
	router := routes.SetupRoutes() // Get router from routes package
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
