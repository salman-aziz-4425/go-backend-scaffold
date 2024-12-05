package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/salman-aziz-4425/Trello-reimagined/pkg/utils"
)

type ContextKey string

const userKey ContextKey = "user"

func ProtectedGuard(next http.Handler) http.Handler {
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
		mapClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		username, ok := mapClaims["username"].(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		id, ok := mapClaims["Id"].(float64)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userKey, struct {
			ID       int
			Username string
		}{
			ID:       int(id),
			Username: username,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
