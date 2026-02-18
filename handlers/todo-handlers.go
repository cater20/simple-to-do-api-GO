package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourname/go-todo-api/models"
	"github.com/yourname/go-todo-api/storage"
)

// CreateTodo creates a new To-Do item
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Assign ID and store in memory
	//todo.ID = fmt.Sprintf("%d", len(storage.Todos)+1)
	//storage.Todos = append(storage.Todos, todo)
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(todo)

	result, err := storage.DB.Exec(
		"INSERT INTO todos (title, completed) VALUES (?, ?)",
		todo.Title, todo.Completed,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	todo.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}

// GetTodos returns all To-Do items
func GetTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := storage.DB.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// GetTodoByID returns a single To-Do by ID
//func GetTodoByID(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)

	//for _, item := range storage.Todos {
	//	if item.ID == params["id"] {
		//	json.NewEncoder(w).Encode(item)
		//	return
	//	}
	//}
	//http.Error(w, "To-Do item not found", http.StatusNotFound)
//}
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var todo models.Todo
	err := storage.DB.QueryRow(
		"SELECT id, title, completed FROM todos WHERE id = ?",
		params["id"],
	).Scan(&todo.ID, &todo.Title, &todo.Completed)

	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo updates a To-Do by ID
//func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)

//	for index, item := range storage.Todos {
		//if item.ID == params["id"] {
			//var updatedTodo models.Todo
			//if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
			//	http.Error(w, "Invalid input", http.StatusBadRequest)
			//	return
		//	}

			//updatedTodo.ID = item.ID
			//storage.Todos[index] = updatedTodo
			//json.NewEncoder(w).Encode(updatedTodo)
			//return
		//}
	//}

	//http.Error(w, "To-Do item not found", http.StatusNotFound)
//}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var updated models.Todo
	json.NewDecoder(r.Body).Decode(&updated)

	_, err := storage.DB.Exec(
		"UPDATE todos SET title = ?, completed = ? WHERE id = ?",
		updated.Title, updated.Completed, params["id"],
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updated.ID = 0
	json.NewEncoder(w).Encode(updated)
}


// DeleteTodo removes a To-Do by ID
//func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)

	//for index, item := range storage.Todos {
		//if item.ID == params["id"] {
			//storage.Todos = append(storage.Todos[:index], storage.Todos[index+1:]...)
			//json.NewEncoder(w).Encode(storage.Todos)
			//return
	//	}
	//}

	//http.Error(w, "To-Do item not found", http.StatusNotFound)
//}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	_, err := storage.DB.Exec(
		"DELETE FROM todos WHERE id = ?",
		params["id"],
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

