package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/salman-aziz-4425/Trello-reimagined/internals/db"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
	todos := []models.Todo{}
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}
	json.NewEncoder(w).Encode(todos)
	w.WriteHeader(http.StatusOK)

}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todoData := models.Todo{}
	json.NewDecoder(r.Body).Decode(&todoData)
	err := db.Pool.QueryRow(context.Background(), "INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id", todoData.Title, todoData.Completed).Scan(&todoData.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"id": todoData.ID})
	w.WriteHeader(http.StatusCreated)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}
