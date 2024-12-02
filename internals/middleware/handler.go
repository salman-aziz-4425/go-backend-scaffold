package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/salman-aziz-4425/Trello-reimagined/pkg/utils"
)

type contextKey string

const usernameKey contextKey = "username"

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if mapClaims, ok := claims.(jwt.MapClaims); ok {
			username, ok := mapClaims["username"].(string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			println("Username:", username)
			ctx := context.WithValue(r.Context(), usernameKey, username)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

}
