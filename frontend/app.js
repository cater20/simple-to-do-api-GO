const apiUrl = "http://localhost:8080/todos";
const form = document.getElementById("todo-form");
const todoList = document.getElementById("todo-list");

// Load todos from API
async function loadTodos() {
    const res = await fetch(apiUrl);
    const todos = await res.json();
    todoList.innerHTML = "";

    todos.forEach(todo => {
        const li = document.createElement("li");
        li.textContent = todo.title;
        li.className = todo.completed ? "completed" : "";

        // Toggle completed on click
        li.addEventListener("click", () => toggleTodo(todo));

        // Delete button
        const delBtn = document.createElement("button");
        delBtn.textContent = "Delete";
        delBtn.addEventListener("click", (e) => {
            e.stopPropagation();
            deleteTodo(todo.id);
        });

        li.appendChild(delBtn);
        todoList.appendChild(li);
    });
}

// Add new todo
form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const title = document.getElementById("title").value;

    await fetch(apiUrl, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, completed: false })
    });

    form.reset();
    loadTodos();
});

// Delete todo
async function deleteTodo(id) {
    await fetch(`${apiUrl}/${id}`, { method: "DELETE" });
    loadTodos();
}

// Toggle completed
async function toggleTodo(todo) {
    await fetch(`${apiUrl}/${todo.id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title: todo.title, completed: !todo.completed })
    });
    loadTodos();
}

// Initial load
loadTodos();
