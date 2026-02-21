# Go To-Do API 📝

🟢 Getting Started with Go – Simple To-Do API (Beginner Friendly Guide)
1️⃣ Title & Objective
Building a RESTful To-Do API in Go – A Beginner’s Guide
🔹 Technology Chosen
Language: Go (Golang)
Router: Gorilla Mux
Database: SQLite

🔹 Why Go?
I chose Go because:

It is simple and beginner-friendly
It is powerful for backend development
It has built-in concurrency
It is widely used in cloud and DevOps tools
It is fast and compiled

As someone transitioning into backend and API development, Go provides a clean syntax and strong performance without unnecessary complexity.

🔹 End Goal
## The goal of this project is to:
✔ Build a RESTful API
✔ Perform CRUD operations (Create, Read, Update, Delete)
## Features
- CRUD operations
- REST API
- SQLite database
- Modular project structure
✔ Connect to a database (SQLite)
✔ Return JSON responses
✔ Run a working backend server on localhost:8080
## Endpoints
- GET /todos
- POST /todos
- GET /todos/{id}
- PUT /todos/{id}
- DELETE /todos/{id}

2️⃣ Quick Summary of the Technology
What is Go?
Go is an open-source programming language developed at Google in 2009.

It is:
Compiled
Statically typed
Fast
Simple and clean

##Where is Go used?
Backend APIs
Microservices
Cloud systems
DevOps tools

##Real-world Example
Docker and Kubernetes are written in Go.
This shows how powerful and scalable Go applications can be.

3️⃣ System Requirements
🖥 Operating System
Windows
Linux
MacOS

🛠 Tools Required
Go (1.20+ recommended)
VS Code (or any code editor)
Git
Terminal / Command Prompt
Postman or curl (for testing)

📦 Required Packages
go get github.com/gorilla/mux
go get github.com/mattn/go-sqlite3

4️⃣ Installation & Setup Instructions
Step 1: Install Go
Download from:
👉 https://go.dev/dl/

Verify installation:
go version

Expected Output:
go version go1.21 windows/amd64

Step 2: Clone the Repository
git clone https://github.com/cater20/simple-to-do-api-GO.git
cd simple-to-do-api-GO

Step 3: Initialize Go Modules
go mod init simple-to-do-api
go mod tidy

Step 4: Run the Server
go run main.go

Expected Output:
Server started at http://localhost:8080

Step 5: Test the API
Check health endpoint:

curl http://localhost:8080/health

Expected Output:
{
  "status": "OK"
}

Create a new Todo:
curl -X POST http://localhost:8080/todos \
-H "Content-Type: application/json" \
-d '{"title":"Learn Go"}'

5️⃣ Minimal Working Example

##What Does This API Do?
Creates To-Do items
Lists all To-Do items
Updates a To-Do
Deletes a To-Do

Example: Creating a Todo (Handler)
func CreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo models.Todo

    // Decode JSON request body
    json.NewDecoder(r.Body).Decode(&todo)

    // Insert into database
    storage.DB.Create(&todo)

    // Return JSON response
    json.NewEncoder(w).Encode(todo)
}
Explanation:

json.NewDecoder reads JSON input
Database saves the todo
json.NewEncoder sends JSON response

Expected JSON Output
{
  "id": 1,
  "title": "Learn Go",
  "completed": false
}

6️⃣ AI Prompt Journal
Prompt 1

I'm not proficient in GO and want to learn GO to Create a simple todo list

Could you help me create a structured learning journey plan with:
- 3-4 distinct learning phases
- Prerequisites for each phase
- 4-5 specific learning steps per phase
- Verification activities for each phase

Prompt 2: 

I want to improve my understanding of CRUDE  in GO .

1. Could you explain this feature with simple examples?
2. Show me 3 practical use cases where this would be valuable
3. Provide a small project idea that would help me practice this feature
4. What common mistakes should I avoid when using this feature?
  
prompt 3 
I want to understand crude in GO just focus on structure and concepts. Could you break down:
1. How IT is implemented 
2. How it works .
3. The key syntax and structures I need to understand
4. Common patterns and best practices

prompt 4 :
I want to build a simple to do list in Go. 
Could you help me:
1. Break this project into small, manageable components
2. Suggest what library/framework components I should use
3. Outline the key files/classes I'll need to create
4. Identify potential challenges I might face coming 

Prompt 5
I've completed my simple todo list project:

https://github.com/cater20/simple-to-do-api-GO.git

Could you:
1. Review the code for GO idioms and best practices
2. Suggest refactoring opportunities to make it more idiomatic
3. Identify any remaining GO patterns I should adjust
4. Recommend next steps to continue improving my GO skills

7️⃣ Common Issues & Fixes
❌ Issue 1: 404 Page Not Found

Cause: Route not registered correctly.
Fix: Ensure:

router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
❌ Issue 2: curl Cannot Connect

Error:

curl: (7) Failed to connect to localhost port 8080

Cause: Server not running
Fix: Run:

go run main.go
❌ Issue 3: Untracked Git Files

Message:

Untracked files: ../README.md

Fix:

git add .
git commit -m "Added README"
git push
8️⃣ References
Official Documentation

Go Docs → https://go.dev/doc/
Gorilla Mux → https://github.com/gorilla/mux
SQLite → https://www.sqlite.org/docs.html

Helpful Resources
Go by Example → https://gobyexample.com/
REST API Design Guide → https://restfulapi.net/

📌 Final Thoughts

This project helped me understand:

✔ How REST APIs work
✔ HTTP methods (GET, POST, PUT, DELETE)
✔ JSON encoding/decoding
✔ Routing in Go
✔ Database integration
✔ Debugging 404 errors

##For a beginner transitioning into backend development, this project builds strong foundations in:

API architecture
Server setup
Database interaction
Clean project structuring

🚀 Future Improvements

Add JWT Authentication
Add Logging Middleware
Add Docker support
Add Unit Tests
Deploy to cloud (Render, Railway, AWS)
   
