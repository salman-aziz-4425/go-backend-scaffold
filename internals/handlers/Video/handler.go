package video

import (
	"encoding/json"
	"net/http"

	"github.com/salman-aziz-4425/Trello-reimagined/internals/dtos"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/middleware"
)

func AddGroup(w http.ResponseWriter, r *http.Request) {
	println("AddGroup")
	user, ok := r.Context().Value(middleware.ContextKey("user")).(struct {
		ID       int
		Username string
	})
	if !ok {
		http.Error(w, "Username not found in context", http.StatusInternalServerError)
		return
	}
	var requestBody struct {
		GroupName string `json:"groupName"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request: " + err.Error()})
		return
	}
	err = AddGroupLogic(dtos.VideoGroupDTO{
		AuthorId:  user.ID,
		GroupName: requestBody.GroupName,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error: " + err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetGroupDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, ok := r.Context().Value(middleware.ContextKey("user")).(struct {
		ID       int
		Username string
	})
	if !ok {
		http.Error(w, "Username not found in context", http.StatusInternalServerError)
		return
	}
	groups, err := GetGroupDetailsLogic(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error: " + err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}
