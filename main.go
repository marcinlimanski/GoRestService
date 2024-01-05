package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Todo struct represents a simple Todo item.
type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var todos []Todo
var currentID = 1

func main() {
	// Create a new router instance from Gorilla Mux.
	r := mux.NewRouter()

	// Define your RESTful API routes.
	r.HandleFunc("/todos", GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	r.HandleFunc("/todos", CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	// Start the HTTP server on port 8080.
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}

// GetTodos returns a list of all todos.
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// GetTodo returns a single todo by ID.
func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, todo := range todos {
		if todo.ID == id {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

// CreateTodo adds a new todo to the list.
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	todo.ID = currentID
	currentID++
	todos = append(todos, todo)

	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo updates an existing todo by ID.
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, todo := range todos {
		if todo.ID == id {
			var updatedTodo Todo
			_ = json.NewDecoder(r.Body).Decode(&updatedTodo)
			updatedTodo.ID = id
			todos[index] = updatedTodo
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

// DeleteTodo removes a todo by ID.
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
