package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourname/go-todo-api/handlers"
)

// SetupRoutes returns a configured router with all API routes
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Base routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("To-Do API is running 🚀"))
	}).Methods("GET")

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}).Methods("GET")

	// To-Do CRUD routes
	router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", handlers.GetTodoByID).Methods("GET")
	router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	return router
}
