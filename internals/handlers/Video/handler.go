package video

import (
	"encoding/json"
	"net/http"

	"github.com/salman-aziz-4425/Trello-reimagined/internals/models"
)

type ContextKey string

const userKey ContextKey = "user"

func AddGroup(w http.ResponseWriter, r *http.Request) {
	println("AddGroup")
	user, ok := r.Context().Value(userKey).(string)
	if !ok {
		http.Error(w, "Username not found in context", http.StatusInternalServerError)
		return
	}
	println("User:", user)
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
