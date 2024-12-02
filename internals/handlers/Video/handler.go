package video

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/salman-aziz-4425/Trello-reimagined/internals/models"
)

type ContextKey string

const usernameKey ContextKey = "username"

func AddGroup(w http.ResponseWriter, r *http.Request) {
	println("AddGroup")
	username, ok := r.Context().Value(usernameKey).(string)
	if !ok {
		http.Error(w, "Username not found in context", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", username)
	var video = models.Video{}

	err := json.NewDecoder(r.Body).Decode(&video)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request: " + err.Error()})
		return
	}

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
