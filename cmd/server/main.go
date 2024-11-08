package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/db"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/routes"
)

func main() {
	db.Init()

	defer db.Pool.Close()
	router := mux.NewRouter()
	routes.SetupRoutes(router)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
