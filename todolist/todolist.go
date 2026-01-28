package todolist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"time"
)

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt int64  `json:"createdAt"`
}

type TodoList struct {
	Items []Todo `json:"items"`
}

var (
	todoTemplate *template.Template
	dataFile     = "data/todos.json"
	todos        = TodoList{Items: []Todo{}}
)

func init() {
	var err error
	todoTemplate, err = template.ParseFiles("templates/todolist.html")
	if err != nil {
		fmt.Printf("Error parsing todolist template: %v\n", err)
	}

	loadTodos()
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/todolist", handleTodoListPage)
	mux.HandleFunc("/todolist/api/todos", handleGetTodos)
	mux.HandleFunc("/todolist/api/add", handleAddTodo)
	mux.HandleFunc("/todolist/api/toggle", handleToggleTodo)
	mux.HandleFunc("/todolist/api/delete", handleDeleteTodo)
}

func handleTodoListPage(w http.ResponseWriter, r *http.Request) {
	todoTemplate.Execute(w, todos.Items)
}

func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos.Items)
}

func handleAddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	newTodo := Todo{
		ID:        time.Now().UnixNano(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now().Unix(),
	}

	todos.Items = append(todos.Items, newTodo)
	saveTodos()

	json.NewEncoder(w).Encode(newTodo)
}

func handleToggleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ID int64 `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range todos.Items {
		if todos.Items[i].ID == req.ID {
			todos.Items[i].Completed = !todos.Items[i].Completed
			saveTodos()
			json.NewEncoder(w).Encode(todos.Items[i])
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ID int64 `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range todos.Items {
		if todos.Items[i].ID == req.ID {
			todos.Items = append(todos.Items[:i], todos.Items[i+1:]...)
			saveTodos()
			fmt.Fprint(w, "{\"success\": true}")
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func saveTodos() {
	os.MkdirAll("data", 0755)
	data, _ := json.MarshalIndent(todos, "", "  ")
	ioutil.WriteFile(dataFile, data, 0644)
}

func loadTodos() {
	os.MkdirAll("data", 0755)

	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			todos = TodoList{Items: []Todo{}}
			return
		}
		fmt.Printf("Error reading todos: %v\n", err)
		todos = TodoList{Items: []Todo{}}
		return
	}

	var loaded TodoList
	err = json.Unmarshal(data, &loaded)
	if err != nil {
		fmt.Printf("Error parsing todos: %v\n", err)
		todos = TodoList{Items: []Todo{}}
		return
	}

	todos = loaded
}
