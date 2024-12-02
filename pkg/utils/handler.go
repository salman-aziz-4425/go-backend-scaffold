package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/config"
)

func CreateToken(username string) (string, error) {
	secretKey := []byte(config.LoadConfig().SecretKey)
	if len(secretKey) == 0 {
		return "", errors.New("invalid secret key")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("JWT signing error: %w", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	fmt.Println("Raw Token String:", tokenString)
	secretKey := config.LoadConfig().SecretKey
	fmt.Println("Secret Key:", secretKey)

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

	fmt.Printf("Token Header: %v\n", token.Header)
	fmt.Printf("Token Claims: %v\n", token.Claims)

	return token.Claims, nil
}
