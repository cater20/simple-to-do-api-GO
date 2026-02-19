package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cater20/go-todo-api/handlers"
	"github.com/cater20/go-todo-api/storage"
	"github.com/cater20/go-todo-api/web"
	"github.com/gorilla/mux"
)

func main() {
	storage.InitDB()

	router := mux.NewRouter()

	// Frontend routes
	router.HandleFunc("/", web.HomeHandler).Methods("GET")
	router.HandleFunc("/add", web.AddTodoHandler).Methods("POST")

	// API routes
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!"))
	}).Methods("GET")
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")
	router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", handlers.GetTodoByID).Methods("GET")
	router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
