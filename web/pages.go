package web

import (
	"html/template"
	"net/http"

	"github.com/cater20/go-todo-api/storage"
)

var tmpl = template.Must(template.ParseGlob("web/templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := storage.DB.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		http.Error(w, "DB Error: "+err.Error(), 500)
		return
	}
	defer rows.Close()

	var todos []struct {
		ID        int
		Title     string
		Completed bool
	}

	for rows.Next() {
		var t struct {
			ID        int
			Title     string
			Completed bool
		}
		rows.Scan(&t.ID, &t.Title, &t.Completed)
		todos = append(todos, t)
	}

	err = tmpl.ExecuteTemplate(w, "home.html", todos)
	if err != nil {
		http.Error(w, "Template Error: "+err.Error(), 500)
		return
	}
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	_, err := storage.DB.Exec("INSERT INTO todos(title, completed) VALUES (?, ?)", title, false)
	if err != nil {
		http.Error(w, "DB Error: "+err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
