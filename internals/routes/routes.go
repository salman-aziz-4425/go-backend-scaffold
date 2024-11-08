package routes

import (
	"github.com/gorilla/mux"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/handlers"
)

func SetupRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	return router
}
