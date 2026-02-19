package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cater20/go-todo-api/routes"
	"github.com/cater20/go-todo-api/storage"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize SQLite DB
	storage.InitDB()

	// Create main router
	router := mux.NewRouter()

	// Setup API routes
	apiRouter := routes.SetupRoutes()        // returns *mux.Router
	router.PathPrefix("/todos").Handler(apiRouter)

	// Serve frontend static files
	fs := http.FileServer(http.Dir("./frontend"))
	router.PathPrefix("/").Handler(fs)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
