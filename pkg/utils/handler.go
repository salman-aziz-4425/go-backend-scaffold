package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/config"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/models"
)

func CreateToken(u models.User) (string, error) {
	secretKey := []byte(config.LoadConfig().SecretKey)
	if len(secretKey) == 0 {
		return "", errors.New("invalid secret key")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"Id":       u.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("JWT signing error: %w", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	secretKey := config.LoadConfig().SecretKey
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println("Error while parsing token:", err)
		return nil, err
	}
	if token == nil || !token.Valid {
		fmt.Println("Token is invalid")
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims type")
	}
	println("Username:", claims)
	return claims, nil
}
