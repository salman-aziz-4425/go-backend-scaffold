package User

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/salman-aziz-4425/Trello-reimagined/internals/dtos"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dtos.UserLoginDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request: " + err.Error()})
		return
	}
	println(user.Username)
	println(user.Password)
	tokenString, err := LoginLogic(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error logging in"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprint(w, "Invalid request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenString, err := RegisterLogic(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error registering")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
