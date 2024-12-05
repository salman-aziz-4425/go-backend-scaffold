package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	chat "github.com/salman-aziz-4425/Trello-reimagined/internals/handlers/Chat"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/handlers/User"
	video "github.com/salman-aziz-4425/Trello-reimagined/internals/handlers/Video"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/middleware"
)

func SetupRoutes(router *mux.Router) *mux.Router {
	manager := chat.NewClientManager()
	go manager.HandleMessages()
	router.HandleFunc("/login", User.Login).Methods("POST")
	router.HandleFunc("/register", User.Register).Methods("POST")
	router.Handle("/group", middleware.ProtectedGuard(http.HandlerFunc(video.AddGroup))).Methods("POST")
	router.Handle("/groupDetails", middleware.ProtectedGuard(http.HandlerFunc(video.GetGroupDetails))).Methods("GET")
	router.HandleFunc("/ws", manager.HandleConnections)

	return router
}
