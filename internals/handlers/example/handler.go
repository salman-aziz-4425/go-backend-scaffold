package example

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/db"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []models.Todo{}
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM todo")
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
	err := db.Pool.QueryRow(context.Background(), "INSERT INTO todo (title, completed) VALUES ($1, $2) RETURNING id", todoData.Title, todoData.Completed).Scan(&todoData.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"id": todoData.ID})
	w.WriteHeader(http.StatusCreated)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := models.Todo{}
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM todo WHERE ID = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		err = rows.Scan(&row.ID, &row.Title, &row.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(row)
	w.WriteHeader(http.StatusOK)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	reqData := models.Todo{}
	json.NewDecoder(r.Body).Decode(&reqData)
	vars := mux.Vars(r)
	id := vars["id"]
	_, err := db.Pool.Exec(context.Background(), "UPDATE todo SET title = $1, completed = $2 WHERE id = $3", reqData.Title, reqData.Completed, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}
